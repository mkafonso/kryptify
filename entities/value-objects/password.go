package valueobjects

import "golang.org/x/crypto/bcrypt"

type Password string

const bcryptCost = 12

func NewPassword(password string) Password {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return Password(hashedPassword)
}

func (p Password) IsPasswordValid(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(password))
	return err == nil
}
