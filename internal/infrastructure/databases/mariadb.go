package databases

import (
	"context"
	"platform/internal/infrastructure/configs"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
)

type MariaDBSession struct {
	*Session
}

func NewMariaDBSession(
	log *zap.Logger,
	cfg *configs.DBConfig,
) (*MariaDBSession, error) {

	session, err := NewSession(
		log,
		cfg,
		mysql.Open,
	)

	if err != nil {
		return nil, err
	}

	return &MariaDBSession{
		Session: session,
	}, nil
}

func (s *MariaDBSession) Version(ctx context.Context) (string, error) {
	var version string

	err := s.
		Gorm(ctx).
		Raw("SELECT version()").
		Scan(&version).
		Error

	return version, err
}
