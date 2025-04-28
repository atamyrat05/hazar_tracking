package utilits

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	securityKey = "qwerty2005"
	tokenTLL    = time.Hour * 5
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int
}

func GenerateToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTLL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	return token.SignedString([]byte(securityKey))
}

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign-in method")
		}
		return []byte(securityKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not type of *tokenClaims")
	}
	return claims.UserId, nil
}
