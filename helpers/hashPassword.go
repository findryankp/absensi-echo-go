package helpers

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "err"
	}

	password = string(bytes)
	return password
}

func CheckPassword(existPassword, requestPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(existPassword), []byte(requestPassword))
	if err != nil {
		return false
	}
	return true
}
