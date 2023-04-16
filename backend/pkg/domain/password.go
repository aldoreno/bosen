package domain

import "golang.org/x/crypto/bcrypt"

type Password string

func (p Password) CheckPasswordHash(plainTextPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(plainTextPassword))
	return err == nil
}
