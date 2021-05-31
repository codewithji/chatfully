package handler

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(pw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}

// func Login(c *fiber.Ctx) error {

// }
