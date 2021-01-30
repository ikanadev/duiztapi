package services

import (
	"time"

	"github.com/gbrlsnchs/jwt/v3"
)

// JWTPayload the struct of jwt data with claims
type JWTPayload struct {
	jwt.Payload
	ID int
}

// GenerateToken generates the claims needed
func GenerateToken(secret string, ID int) (string, error) {
	now := time.Now()
	payload := JWTPayload{
		Payload: jwt.Payload{
			Issuer:         "Duizt App",
			IssuedAt:       jwt.NumericDate(now),
			ExpirationTime: jwt.NumericDate(now.Add(7 * 24 * time.Hour)),
		},
		ID: ID,
	}
	token, err := jwt.Sign(payload, jwt.NewHS256([]byte(secret)))
	if err != nil {
		return "", err
	}
	return string(token), nil
}

// CheckToken takes a token string and returns the ID claim
func CheckToken(secret string, token string) (int, error) {
	jwtPayload := JWTPayload{}
	_, err := jwt.Verify([]byte(token), jwt.NewHS256([]byte(secret)), &jwtPayload)
	if err != nil {
		return 0, err
	}
	return jwtPayload.ID, err
}
