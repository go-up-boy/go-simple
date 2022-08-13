package types

import (
	"database/sql"
	"gorm.io/gorm"
)

type ConnectionStruct struct {
	Driver string
	DB *gorm.DB
	SqlDB *sql.DB
}