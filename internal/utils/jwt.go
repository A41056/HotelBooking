package utils

import (
	"errors"
	"time"

	_const "dev.longnt1.git/aessment-hotel-booking.git/internal/const"

	"github.com/dgrijalva/jwt-go"
)

var JWTSecretKey = []byte("RdPtxsF0GY0dQ2mFm25v2giMaPzo4DXf")

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) (string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", errors.New(_const.ErrInvalidTokenSignature)
		}
		return "", err
	}

	if !token.Valid {
		return "", errors.New(_const.ErrInvalidToken)
	}

	return claims.UserID, nil
}
