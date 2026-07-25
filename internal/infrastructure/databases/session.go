package databases

import (
	"context"
	"errors"
	"fmt"
	"platform/internal/infrastructure/configs"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Session struct {
	log *zap.Logger
	db  *gorm.DB
}

func NewSession(
	log *zap.Logger,
	cfg *configs.DBConfig,
	open OpenFunc,
) (*Session, error) {

	db, err := NewGorm(
		log,
		cfg,
		open,
	)

	if err != nil {
		return nil, err
	}

	return &Session{
		log: log,
		db:  db,
	}, nil
}

// Gorm повертає поточну сесію z прив'язаним context.
// Якщо в context вже є транзакція — повертається вона.
func (s *Session) Gorm(ctx context.Context) *gorm.DB {
	if tx, ok := transactionFromContext(ctx); ok {
		return tx
	}

	return s.db.WithContext(ctx)
}

// WithTransaction запускає callback всередині транзакції.
func (s *Session) WithTransaction(
	ctx context.Context,
	fn func(context.Context) error,
) error {
	if _, ok := transactionFromContext(ctx); ok {
		return fn(ctx)
	}

	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("begin transaction: %w", tx.Error)
	}

	s.log.Debug("transaction started")

	defer func() {
		if r := recover(); r != nil {
			if err := tx.Rollback().Error; err != nil {
				s.log.Error(
					"rollback after panic failed",
					zap.Error(err),
				)
			}
			s.log.Error("transaction rolled back due to panic", zap.Any("panic", r))
			panic(r)
		}
	}()

	ctx = withTransaction(ctx, tx)

	if err := fn(ctx); err != nil {
		if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
			s.log.Error("rollback transaction failed", zap.Error(rollbackErr))
			return errors.Join(err, rollbackErr)
		}

		s.log.Debug("transaction rollback", zap.Error(err))
		return err
	}

	if err := tx.Commit().Error; err != nil {
		s.log.Error("commit transaction failed", zap.Error(err))
		return fmt.Errorf("commit transaction: %w", err)
	}

	s.log.Debug("transaction commit")

	return nil
}
