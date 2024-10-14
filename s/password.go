package util

import (
	dto "github.com/srv-cashpay/auth/dto/auth"

	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func EncryptPassword(req *dto.SignupRequest) (err error) {
	hashedPassword, err := GenerateFromPassword(req.Whatsapp)

	if err != nil {
		return err
	}
	req.Whatsapp = string(hashedPassword)
	return nil
}
