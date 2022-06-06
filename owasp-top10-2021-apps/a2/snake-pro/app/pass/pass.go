package pass

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// CheckPass checks a password
func CheckPass(truePassword, attemptPassword string) bool {
	return truePassword == attemptPassword
}

func GeneratePassHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	fmt.Println(string(bytes))
	return string(bytes), err
}

func CheckPassHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(password, hash, err)

	return err == nil
}
