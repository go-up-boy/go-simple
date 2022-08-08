package hash

import (
	"go-simple/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func BcryptHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	logger.LogIf(err)

	return string(bytes)
}

func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func BcryptIsHashed(str string) bool {
	return len(str) == 60
}