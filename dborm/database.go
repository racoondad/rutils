package dborm

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// CreateDatabase 创建数据库
func CreateDatabase(dbType, username, password, host, dbName string, port uint) error {
	switch dbType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/mysql?charset=utf8mb4&parseTime=True&loc=Local",
			username,
			password,
			host,
			port,
		)
		conn, err := sql.Open("mysql", dsn)
		if err != nil {
			fmt.Println("自动创建数据库失败")
		}
		defer conn.Close()
		_, err = conn.Exec(fmt.Sprintf("create database %s;", dbName))
		return err
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%d sslmode=disable TimeZone=Asia/Shanghai",
			host,
			username,
			password,
			port,
		)
		conn, err := sql.Open("postgres", dsn)
		if err != nil {
			fmt.Println("自动创建数据库失败")
		}
		defer conn.Close()
		_, err = conn.Exec(fmt.Sprintf("create database %s;", dbName))
		return err
	case "sqlserver":
		dsn := fmt.Sprintf("server=%s;port%d;database=master;user id=%s;password=%s", host, port, username, password)
		conn, err := sql.Open("mssql", dsn)
		if err != nil {
			fmt.Println("自动创建数据库失败")
		}
		defer conn.Close()
		_, err = conn.Exec(fmt.Sprintf("create database %s;", dbName))
		return err
	default:
		return nil
	}
}

// DropDatabase 删除数据库(需要登陆管理库)
func DropDatabase(dbType, username, password, host, dbName string, port uint) error {
	switch dbType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/mysql?charset=uutf8mb4&parseTime=True&loc=Local",
			username,
			password,
			host,
			port,
		)
		conn, err := sql.Open("mysql", dsn)
		if err != nil {
			fmt.Println("自动创建数据库失败")
		}
		defer conn.Close()
		_, err = conn.Exec(fmt.Sprintf("drop database %s;", dbName))
		return err
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%d sslmode=disable TimeZone=Asia/Shanghai",
			host,
			username,
			password,
			port,
		)
		conn, err := sql.Open("postgres", dsn)
		if err != nil {
			fmt.Println("自动创建数据库失败")
		}
		defer conn.Close()
		_, err = conn.Exec(fmt.Sprintf("drop database %s;", dbName))
		return err
	case "sqlserver":
		dsn := fmt.Sprintf("server=%s;port%d;database=master;user id=%s;password=%s", host, port, username, password)
		conn, err := sql.Open("mssql", dsn)
		if err != nil {
			fmt.Println("自动创建数据库失败")
		}
		defer conn.Close()
		_, err = conn.Exec(fmt.Sprintf("drop database %s;", dbName))
		return err
	default:
		return nil
	}
}
