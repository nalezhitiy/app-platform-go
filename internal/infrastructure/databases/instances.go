package databases

import (
	"platform/internal/infrastructure/configs"

	"go.uber.org/zap"
)

// POSTGRES Sessions
type MainPostgresSession struct {
	*PostgresSession
}

func NewMainPostgresSession(
	log *zap.Logger,
	cfg *configs.PlatformDBConfig,
) (*MainPostgresSession, error) {

	log = log.With(
		zap.String("database", "main-postgres"),
	)

	session, err := NewPostgresSession(log, &cfg.Postgres)
	if err != nil {
		return nil, err
	}

	return &MainPostgresSession{
		PostgresSession: session,
	}, nil
}

//MARIADB Ssession

type MainMariaDBSession struct {
	*MariaDBSession
}

func NewMainMariaDBSession(
	log *zap.Logger,
	cfg *configs.PlatformDBConfig,
) (*MainMariaDBSession, error) {

	log = log.With(
		zap.String("database", "mariadb"),
	)

	session, err := NewMariaDBSession(log, &cfg.MariaDB)
	if err != nil {
		return nil, err
	}

	return &MainMariaDBSession{
		MariaDBSession: session,
	}, nil
}

// CLICKHOUSE Session

type MainClickhouseSession struct {
	*ClickHouseSession
}

func NewMainClickhouseSession(log *zap.Logger,
	cfg *configs.PlatformDBConfig) (*MainClickhouseSession, error) {

	log = log.With(
		zap.String("database", "clickhouse"),
	)

	session, err := NewClickHouseSession(log, &cfg.ClickHouse)
	if err != nil {
		return nil, err
	}

	return &MainClickhouseSession{
		ClickHouseSession: session,
	}, nil

}
