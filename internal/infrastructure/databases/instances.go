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

	dbCfg, err := cfg.Get("postgres-main")
	if err != nil {
		return nil, err
	}

	log = log.With(
		zap.String("database", dbCfg.Name),
	)

	session, err := NewPostgresSession(log, dbCfg)
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

	dbCfg, err := cfg.Get("mariadb-main")
	if err != nil {
		return nil, err
	}

	log = log.With(
		zap.String("database", dbCfg.Name),
	)

	session, err := NewMariaDBSession(log, dbCfg)
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

	dbCfg, err := cfg.Get("clickhouse-main")
	if err != nil {
		return nil, err
	}

	log = log.With(
		zap.String("database", dbCfg.Name),
	)

	session, err := NewClickHouseSession(log, dbCfg)
	if err != nil {
		return nil, err
	}

	return &MainClickhouseSession{
		ClickHouseSession: session,
	}, nil

}
