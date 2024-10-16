package util

import (
	"math/rand"
	"strings"
	"time"

	dto "github.com/srv-cashpay/auth/dto/auth"

	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateFromWhatsapp(password string) ([]byte, error) {
	return bcrypt.GenerateFromWhatsapp([]byte(password), bcrypt.DefaultCost)
}

func VerifyWhatsapp(hashedPassword, password string) error {
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

func GenerateRandomNumeric(length int) string {
	const chars = "0123456789"

	var result strings.Builder
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		result.WriteRune(rune(chars[rand.Intn(len(chars))]))
	}

	return result.String()
}
