// Package user 模型
package user

import (
    "go-simple/app/models"
    "go-simple/pkg/database"
)

type User struct {
    models.BaseModel

    // Put fields in here
    Name     string `gorm:"type:varchar(255);not null;index"`
    Username string `gorm:"type:varchar(255);not null;index"`
    Email    string `gorm:"type:varchar(255);index;default:null"`
    Phone    string `gorm:"type:varchar(20);index;default:null"`
    Password string `gorm:"type:varchar(255)" json:"-"`

    models.CommonTimestampsField
}

func (user *User) Create() {
    database.DB.Create(&user)
}

func (user *User) Save() (rowsAffected int64) {
    result := database.DB.Save(&user)
    return result.RowsAffected
}

func (user *User) Delete() (rowsAffected int64) {
    result := database.DB.Delete(&user)
    return result.RowsAffected
}