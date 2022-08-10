package bootstrap

import (
	"errors"
	"fmt"
	"go-simple/pkg/config"
	"go-simple/pkg/database"
	"go-simple/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

func SetupDB() {

	var dbConfig gorm.Dialector

	switch config.Get("database.drive") {
	case "mysql":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.mysql.username"),
			config.Get("database.mysql.password"),
			config.Get("database.mysql.host"),
			config.Get("database.mysql.port"),
			config.Get("database.mysql.database"),
			config.Get("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
		break
	case "sqlite":
		database := config.Get("database.sqlite.database")
		dbConfig = sqlite.Open(database)
		break
	default:
		panic(errors.New("database connection not supported"))
	}
	database.Connect(dbConfig, logger.NewGormLogger())
	// Set the maximum number of connections
	database.SqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// Set the maximum number of idle connections
	database.SqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// Set the expiration time of each link
	database.SqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
}
