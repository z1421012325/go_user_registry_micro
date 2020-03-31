package utils

import (
	"golang.org/x/crypto/bcrypt"
)


func Encryption(out string,cost int) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(out), cost)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func CheckEncryption(now,input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(now), []byte(input))
	return err == nil
}