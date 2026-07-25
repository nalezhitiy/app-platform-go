package databases

import (
	"platform/internal/infrastructure/configs"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func InjectDI() fx.Option {
	return fx.Options(

		fx.Provide(

			func(
				log *zap.Logger,
				cfg *configs.Config,
			) (*PostgresSession, error) {

				log = log.With(
					zap.String("database", "main-postgres"),
				)

				return NewPostgresSession(
					log,
					&cfg.Databases.Postgres,
				)
			},

			func(
				log *zap.Logger,
				cfg *configs.Config,
			) (*MariaDBSession, error) {

				log = log.With(
					zap.String("database", "mariadb"),
				)

				return NewMariaDBSession(
					log,
					&cfg.Databases.MariaDB,
				)
			},

			func(
				log *zap.Logger,
				cfg *configs.Config,
			) (*ClickHouseSession, error) {

				log = log.With(
					zap.String("database", "clickhouse"),
				)

				return NewClickHouseSession(
					log,
					&cfg.Databases.ClickHouse,
				)
			},
		),
	)
}
