package services

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"e-commerce/internal/entities"
	"e-commerce/internal/storage/repositories"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repository repositories.AuthCrudRepository
}

func NewAuthService(repo repositories.AuthCrudRepository) *AuthService {
	return &AuthService{repository: repo}
}

func (s *AuthService) Registration(user entities.User) (string, error) {
	user.Password = generatePasswordHash(user.Password)

	_, err := s.repository.CreateUser(user)
	if err != nil {
		return "User registration failed", nil
	}

	return "User registration was successful", nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repository.ReadUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}
