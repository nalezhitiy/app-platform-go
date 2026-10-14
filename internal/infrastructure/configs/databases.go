package configs

import (
	"fmt"
	"platform/internal/utils"
	"time"

	"gorm.io/gorm"
)

type PlatformDBConfig map[string]DBConfig

func (c PlatformDBConfig) Get(name string) (*DBConfig, error) {
	cfg, ok := c[name]
	if !ok {
		return nil, fmt.Errorf("database config %q not found", name)
	}

	return &cfg, nil
}

type DBConfig struct {
	Name       string               `yaml:"name"`
	Driver     string               `yaml:"driver"`
	PrimaryDSN utils.SecretString   `yaml:"primary_dsn"`
	Replicas   []utils.SecretString `yaml:"replicas_dsn"`
	Debug      bool                 `yaml:"debug"`
	Pool       PoolConfig           `yaml:"pool"`
	Orm        OrmConfig            `yaml:"orm"`
}

type PoolConfig struct {
	Min      int           `yaml:"min"`
	Max      int           `yaml:"max"`
	Lifetime time.Duration `yaml:"lifetime"` // Завдяки yaml v3 "900s" автоматично розпарситься в time.Duration
}

type OrmConfig struct {
	SkipDefaultTransaction bool `yaml:"skip_default_transaction"`
	FullSaveAssociations   bool `yaml:"full_save_associations"`
	PrepareStmt            bool `yaml:"prepare_stmt"`
	CreateBatchSize        int  `yaml:"create_batch_size"`
	TranslateError         bool `yaml:"translate_error"`
}

func (c *OrmConfig) ToGorm() *gorm.Config {
	return &gorm.Config{
		// configurable from environment / yaml
		SkipDefaultTransaction: c.SkipDefaultTransaction,
		FullSaveAssociations:   c.FullSaveAssociations,
		PrepareStmt:            c.PrepareStmt,
		CreateBatchSize:        c.CreateBatchSize,
		TranslateError:         c.TranslateError,

		// defaults for now
		DisableNestedTransaction: false,
		AllowGlobalUpdate:        false,
		QueryFields:              false,
		PropagateUnscoped:        false,
	}
}
