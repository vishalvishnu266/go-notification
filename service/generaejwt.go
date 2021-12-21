package service

import (
	"os"
	"time"

	"github.com/cristalhq/jwt/v3"
)

func GenerateJwt(user string) (string, error) {

	key := os.Getenv("CENTRI_HMAC")
	signer, _ := jwt.NewSignerHS(jwt.HS256, []byte(key))
	builder := jwt.NewBuilder(signer)
	token, err := builder.Build(jwt.StandardClaims{
		Subject:   user,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1111111) * time.Second)),
	})
	if err != nil {
		return "", err
	}
	return token.String(), nil
}
