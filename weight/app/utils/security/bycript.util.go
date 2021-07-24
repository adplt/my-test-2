package security

import (
	"crypto/sha256"
	configs "project/app/configs"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
)

const (
	iterations = 1e4
)

func ComparePasswordBycript(hased string, password string) bool {
	// Comparing the password with the hash
	hashedByte := []byte(hased)
	hashedPassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(hashedByte, hashedPassword)
	if err != nil {
		configs.Throw(err)
	}
	return true
}

func HashPasswordBycript(pass string) string {
	password := []byte(pass)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func GenerateHashSha256(pw, salt []byte) []byte {
	ret := make([]byte, len(salt))
	copy(ret, salt)
	return append(ret, pbkdf2.Key(pw, salt, iterations, sha256.Size, sha256.New)...)
}
