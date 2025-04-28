package service

import (
	"errors"
	"fmt"
	"hazar_tracking/internal/model"
	"hazar_tracking/internal/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	securityKey = "qwerty"
	tokenTLL    = 24 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Create(user model.User) (int, error) {
	if user.Password != user.Password_confirm {
		return 0, errors.New("password not confirmed")
	}
	return s.repo.Create(user)
}

func (s *AuthService) GenerateToken(input model.SignIn) (string, error) {
	user, err := s.repo.Get(input.Email, input.Password)
	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTLL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(securityKey))
}

func (l *AuthService) ParseToken(accessToken string) (int, error) {
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
