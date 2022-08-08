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

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(u.Password) {
		u.Password = hash.BcryptHash(u.Password)
	}
	return nil
}