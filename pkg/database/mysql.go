package database

import (
	"fmt"
	"github.com/spf13/cast"
	"go-simple/pkg/logger"
	"go-simple/types"
	"gorm.io/driver/mysql"
	"time"
)

func NewMysqlConnection(config map[string]string) types.ConnectionStruct {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		config["username"],
		config["password"],
		config["host"],
		config["port"],
		config["database"],
		config["charset"],
	)
	dbConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	MysqlDB, MysqlSqlDB := ConnectService(dbConfig, logger.NewGormLogger())
	// Set the maximum number of connections
	SqlDB.SetMaxOpenConns(cast.ToInt(config["max_open_connections"]))
	// Set the maximum number of idle connections
	SqlDB.SetMaxIdleConns(cast.ToInt(config["max_idle_connections"]))
	// Set the expiration time of each link
	SqlDB.SetConnMaxLifetime(time.Duration(cast.ToInt(config["max_life_seconds"])) * time.Second)

	return types.ConnectionStruct{
		Driver: "mysql",
		DB: MysqlDB,
		SqlDB: MysqlSqlDB,
	}
}