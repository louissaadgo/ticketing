package token

import (
	"fmt"
	"os"
	"time"

	"github.com/louissaadgo/ticketing/auth/src/models"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

func GeneratePasetoToken(email string) (string, error) {

	payload := models.PasetoPayload{
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(24 * time.Hour),
	}

	key := os.Getenv("PASETO_KEY")

	if len(key) != chacha20poly1305.KeySize {
		return "", fmt.Errorf("key error")
	}

	paseto := paseto.NewV2()

	return paseto.Encrypt([]byte(key), payload, nil)

}

func VerifyPasetoToken(token string) (models.PasetoPayload, bool) {

	payload := models.PasetoPayload{}

	key := os.Getenv("PASETO_KEY")

	paseto := paseto.NewV2()

	err := paseto.Decrypt(token, []byte(key), &payload, nil)
	if err != nil {
		return payload, false
	}

	if !time.Now().Before(payload.ExpiredAt) {
		return payload, false
	}

	return payload, true
}
