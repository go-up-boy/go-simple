package database

import (
	"github.com/spf13/cast"
	"go-simple/pkg/logger"
	"go-simple/types"
	"gorm.io/driver/sqlite"
	"time"
)

func NewSqliteConnection(config map[string]string) types.ConnectionStruct {
	dbConfig := sqlite.Open(config["database"])
	SqliteDB, SqliteSqlDB := ConnectService(dbConfig, logger.NewGormLogger())
	// Set the maximum number of connections
	SqlDB.SetMaxOpenConns(cast.ToInt(config["max_open_connections"]))
	// Set the maximum number of idle connections
	SqlDB.SetMaxIdleConns(cast.ToInt(config["max_idle_connections"]))
	// Set the expiration time of each link
	SqlDB.SetConnMaxLifetime(time.Duration(cast.ToInt(config["max_life_seconds"])) * time.Second)

	return types.ConnectionStruct{
		Driver: "sqlite",
		DB: SqliteDB,
		SqlDB: SqliteSqlDB,
	}
}
