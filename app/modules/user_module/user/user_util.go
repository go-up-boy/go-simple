package user

import (
    "go-simple/pkg/database"
    "go-simple/pkg/hash"
)

func Get(idstr string) (user User) {
    database.DB.Where("id", idstr).First(&user)
    return
}

func GetBy(field, value string) (user User) {
    database.DB.Where("? = ?", field, value).First(&user)
    return
}

func All() (users []User) {
    database.DB.Find(&users)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(User{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
    var count int64
    database.DB.Model(User{}).Where("email = ?", email).Count(&count)
    return count > 0
}

// IsPhoneExist 判断 Email 已被注册
func IsPhoneExist(phone string) bool {
    var count int64
    database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
    return count > 0
}
// ComparePassword 密码是否正确
func (u User) ComparePassword(_password string) bool {
    return hash.BcryptCheck(_password, u.Password)
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (userModel User) {
    database.DB.Where("phone = ?", phone).First(&userModel)
    return
}

func GetByMulti(loginId string) (userModel User) {
    database.DB.
        Where("username = ?", loginId).
        Or("email = ?", loginId).
        Or("phone = ?", loginId).
        First(&userModel)
    return
}