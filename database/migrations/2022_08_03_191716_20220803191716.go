package migrations

import (
    "database/sql"
    "go-simple/globals"
    "go-simple/pkg/migrate"
    "go-simple/types"
    "gorm.io/gorm"
    "time"
)

func CreateUserTable() {
    var connection *types.ConnectionStruct = &globals.GlobalService.Mysql

    type User struct {
        ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`

        Name     string `gorm:"type:varchar(255);not null;index"`
        Username string `gorm:"type:varchar(255);not null;index"`
        Email    string `gorm:"type:varchar(255);index;default:null"`
        Phone    string `gorm:"type:varchar(20);index;default:null"`
        Password string `gorm:"type:varchar(255)"`

        CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
        UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at,omitempty"`
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&User{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&User{})
    }
    migrate.Add(connection, "2022_08_03_191716_20220803191716", up, down)
}