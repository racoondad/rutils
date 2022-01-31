package dborm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func GormConfig(prefix string, logMode bool) *gorm.Config {

	if logMode {
		return &gorm.Config{
			// DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
				TablePrefix:   prefix,
			},
			Logger: logger.Default.LogMode(logger.Info),
		}
	}
	return &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   prefix,
		},
	}

}
