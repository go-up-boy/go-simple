package user

import (
	"go-simple/pkg/hash"
	"gorm.io/gorm"
)

// func (user *User) BeforeSave(tx *gorm.DB) (err error) {}
// func (user *User) BeforeCreate(tx *gorm.DB) (err error) {}
// func (user *User) AfterCreate(tx *gorm.DB) (err error) {}
// func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (user *User) AfterUpdate(tx *gorm.DB) (err error) {}
// func (user *User) AfterSave(tx *gorm.DB) (err error) {}
// func (user *User) BeforeDelete(tx *gorm.DB) (err error) {}
// func (user *User) AfterDelete(tx *gorm.DB) (err error) {}
// func (user *User) AfterFind(tx *gorm.DB) (err error) {}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(user.Password) {
		user.Password = hash.BcryptHash(user.Password)
	}
	return nil
}