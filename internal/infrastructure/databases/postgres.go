package databases

import (
	"context"
	"platform/internal/infrastructure/configs"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
)

type PostgresSession struct {
	*Session
}

func NewPostgresSession(
	log *zap.Logger,
	cfg *configs.DBConfig,
) (*PostgresSession, error) {

	session, err := NewSession(
		log,
		cfg,
		postgres.Open,
	)

	if err != nil {
		return nil, err
	}

	return &PostgresSession{
		Session: session,
	}, nil
}

func (p *PostgresSession) Version(ctx context.Context) (string, error) {

	var version string

	err := p.
		Gorm(ctx).
		Raw("select version()").
		Scan(&version).
		Error

	return version, err
}
