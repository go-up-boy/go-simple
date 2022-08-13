// Package user 模型
package user

import (
    "github.com/spf13/cast"
    "go-simple/globals"
    "go-simple/types"
    "time"
)
// database connection
var connection *types.ConnectionStruct = &globals.GlobalService.Sqlite

type User struct {
    ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`

    // Put fields in here
    Name     string `gorm:"type:varchar(255);not null;index"`
    Username string `gorm:"type:varchar(255);not null;index"`
    Email    string `gorm:"type:varchar(255);index;default:null"`
    Phone    string `gorm:"type:varchar(20);index;default:null"`
    Password string `gorm:"type:varchar(255)" json:"-"`

    CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
    UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at,omitempty"`
}

func NewUserConnection() *types.ConnectionStruct {
    return connection
}

func (user *User) Create() {
    connection.DB.Create(&user)
}

func (user *User) Save() (rowsAffected int64) {
    result := connection.DB.Save(&user)
    return result.RowsAffected
}

func (user *User) Delete() (rowsAffected int64) {
    result := connection.DB.Delete(&user)
    return result.RowsAffected
}

func (user User) GetStringID() string {
    return cast.ToString(user.ID)
}