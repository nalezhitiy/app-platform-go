package usecases

import (
	"context"
	"platform/internal/infrastructure/databases"
)

type VersionDBUseCase struct {
	postgres   *databases.PostgresSession
	mariadb    *databases.MariaDBSession
	clickhouse *databases.ClickHouseSession
}

func NewVersionDBUseCase(postgres *databases.MainPostgresSession, mariadb *databases.MainMariaDBSession, clickhouse *databases.MainClickhouseSession) *VersionDBUseCase {
	return &VersionDBUseCase{
		postgres:   postgres.PostgresSession,
		mariadb:    mariadb.MariaDBSession,
		clickhouse: clickhouse.ClickHouseSession,
	}
}

type DatabaseVersion struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (uc *VersionDBUseCase) GetVersion() ([]DatabaseVersion, error) {
	var result []DatabaseVersion
	ctx := context.Background()
	if version, err := uc.postgres.Version(ctx); err == nil {
		result = append(result, DatabaseVersion{
			Name:    "postgres",
			Version: version,
		})
	}

	if version, err := uc.mariadb.Version(ctx); err == nil {
		result = append(result, DatabaseVersion{
			Name:    "mariadb",
			Version: version,
		})
	}

	if version, err := uc.clickhouse.Version(ctx); err == nil {
		result = append(result, DatabaseVersion{
			Name:    "clickhouse",
			Version: version,
		})
	}
	return result, nil
}
