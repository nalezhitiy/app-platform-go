package databases

import (
	"context"
	"platform/internal/infrastructure/configs"

	"go.uber.org/zap"
	"gorm.io/driver/clickhouse"
)

type ClickHouseSession struct {
	*Session
}

func NewClickHouseSession(
	log *zap.Logger,
	cfg *configs.DBConfig,
) (*ClickHouseSession, error) {

	session, err := NewSession(
		log,
		cfg,
		clickhouse.Open,
	)

	if err != nil {
		return nil, err
	}

	return &ClickHouseSession{
		Session: session,
	}, nil
}

func (s *ClickHouseSession) Version(ctx context.Context) (string, error) {
	var version string

	err := s.
		Gorm(ctx).
		Raw("SELECT version()").
		Scan(&version).
		Error

	return version, err
}
