package dborm

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"

	"gorm.io/gorm"
)

// GormSqlServer 初始化SqlServer数据库
func GormSqlServer(dsn, prefix string, maxIdle, maxOpen int, defaultStringSize int, logMode bool) (*gorm.DB, error) {
	sqlserverConfig := sqlserver.Config{
		DSN:               dsn,
		DefaultStringSize: defaultStringSize,
	}
	if db, err := gorm.Open(sqlserver.New(sqlserverConfig), GormConfig(prefix, logMode)); err != nil {
		return nil, errors.New("连接数据库失败")
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(maxIdle)
		sqlDB.SetMaxOpenConns(maxOpen)
		return db, nil
	}
}

// GormPostgreSql 初始化PostgreSql数据库
func GormPostgreSql(dsn, prefix string, maxIdle, maxOpen int, logMode bool) (*gorm.DB, error) {
	postgresConfig := postgres.Config{
		DSN:                  dsn,   // DSN data source name
		PreferSimpleProtocol: false, // 禁用隐式 prepared statement
	}
	if db, err := gorm.Open(postgres.New(postgresConfig), GormConfig(prefix, logMode)); err != nil {
		return nil, errors.New("连接数据库失败")
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(maxIdle)
		sqlDB.SetMaxOpenConns(maxOpen)
		return db, nil
	}
}

// GormMysql 初始化PostgreSql数据库
func GormMysql(dsn, prefix string, maxIdle, maxOpen int, defaultStringSize uint, logMode bool) (*gorm.DB, error) {
	mysqlConfig := mysql.Config{
		DSN:                       dsn,               // DSN data source name
		DefaultStringSize:         defaultStringSize, // string 类型字段的默认长度
		DisableDatetimePrecision:  true,              // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,              // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,              // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,             // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), GormConfig(prefix, logMode)); err != nil {
		return nil, errors.New("连接数据库失败")
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(maxIdle)
		sqlDB.SetMaxOpenConns(maxOpen)
		return db, nil
	}
}
