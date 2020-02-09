package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var secret = os.Getenv("JWT_SECRET")

type AuthClaim struct {
	userId UserID
	*jwt.StandardClaims
}
type UserID int

func Token(userID UserID) (string, error) {
	nowUnix := time.Now().Unix()
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	claims := &AuthClaim{
		userID,
		&jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: nowUnix + 60*60,
			Id:        uuid.String(),
			IssuedAt:  nowUnix,
			Issuer:    "",
			NotBefore: nowUnix - 5,
			Subject:   "access_token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Parse(tokenString string) (UserID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return -1, err
	}
	if !token.Valid {
		return -1, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(AuthClaim)
	if !ok {
		return -1, fmt.Errorf("not ok %v", token.Claims)
	}

	return claims.userId, nil
}
