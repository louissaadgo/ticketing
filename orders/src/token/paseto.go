package token

import (
	"os"
	"time"

	"github.com/louissaadgo/ticketing/orders/src/models"
	"github.com/o1egl/paseto"
)

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
