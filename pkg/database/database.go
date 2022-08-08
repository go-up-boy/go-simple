package database

import (
	"database/sql"
	"errors"
	"fmt"
	"go-simple/pkg/config"
	"gorm.io/gorm"
)
import gormlogger "gorm.io/gorm/logger"

var DB *gorm.DB
var SqlDB *sql.DB

func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	SqlDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return
}

func DeleteAllTables() error {
	var err error
	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMySQLTables()
	case "sqlite":
		err = deleteAllSqliteTables()
	default:
		panic(errors.New("database connection not supported"))
	}
	return err
}

func deleteAllSqliteTables() error {
	var tables []string
	err := DB.Raw("SELECT name FROM sqlite_master Where type = 'table'").Pluck("name", &tables).Error
	if err != nil {
		return err
	}
	// Delete ALL TABLES
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteMySQLTables() error {
	dbname := CurrentDatabase()
	var tables []string
	err := DB.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).
		Error
	if err != nil {
		return err
	}
	// Close Foreign Key
	DB.Exec("Set foreign_key_checks = 0;")
	// Delete ALL TABLES
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}
	// Open Foreign Key
	DB.Exec("SET foreign_key_checks = 1;")
	return nil
}