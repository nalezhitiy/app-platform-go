package databases

import (
	"fmt"
	"platform/internal/infrastructure/logger"

	"platform/internal/infrastructure/configs"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type OpenFunc func(string) gorm.Dialector

func NewGorm(
	log *zap.Logger,
	cfg *configs.DBConfig,
	open OpenFunc,
) (*gorm.DB, error) {

	primary := open(string(cfg.PrimaryDSN))

	gormCfg := cfg.Orm.ToGorm()
	// set custom logger
	gormCfg.Logger = logger.NewGormLogger(log, cfg.Debug)

	db, err := gorm.Open(primary, gormCfg)

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.Pool.Min)
	sqlDB.SetMaxOpenConns(cfg.Pool.Max)
	sqlDB.SetConnMaxLifetime(cfg.Pool.Lifetime)

	replicas := make([]gorm.Dialector, 0, len(cfg.Replicas))

	for _, dsn := range cfg.Replicas {

		if dsn == "" {
			continue
		}

		replicas = append(
			replicas,
			open(string(dsn)),
		)
	}

	if len(replicas) > 0 {

		if err := db.Use(
			dbresolver.Register(
				dbresolver.Config{
					Sources: []gorm.Dialector{
						primary,
					},
					Replicas: replicas,
				},
			),
		); err != nil {
			return nil, fmt.Errorf("register dbresolver: %w", err)
		}

		log.Info("GORM initialized", zap.Int("replicas", len(replicas)))
	} else {
		log.Info("GORM initialized", zap.String("mode", "single"))
	}

	if cfg.Debug {
		db = db.Debug()
	}

	return db, nil
}
