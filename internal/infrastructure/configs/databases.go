package configs

import (
	"platform/internal/utils"
	"time"
)

type DatabaseRootConfig struct {
	Database PlatformDBConfig `yaml:"database"`
}

type PlatformDBConfig struct {
	Postgres   DatabaseConfig `yaml:"postgres"`
	MariaDB    DatabaseConfig `yaml:"mariadb"`
	ClickHouse DatabaseConfig `yaml:"clickhouse"`
}

// Спільні параметри для пулу та ORM, які ти використовуєш
type PoolConfig struct {
	Min      int           `yaml:"min"`
	Max      int           `yaml:"max"`
	Lifetime time.Duration `yaml:"lifetime"` // Завдяки yaml v3 "900s" автоматично розпарситься в time.Duration
}

type OrmConfig struct {
	SkipDefaultTransaction bool `yaml:"skip_default_transaction"`
	FullSaveAssociations   bool `yaml:"full_save_associations"`
}

// ------ Специфічні конфіги для кожної БД ------

type DatabaseConfig struct {
	DSN        utils.SecretString `yaml:"dsn"`
	ReplicaDSN utils.SecretString `yaml:"replica_dsn"`
	Debug      bool               `yaml:"debug"`
	Pool       PoolConfig         `yaml:"pool"`
	Orm        OrmConfig          `yaml:"orm"`
}
