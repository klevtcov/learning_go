package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/klevtcov/todo"
	"github.com/klevtcov/todo/pkg/repository"
)

const (
	salt       = "sdlkfjgne56welkjwbfg63er"
	signingKey = "jkasdfajs6523tfrgksddf"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Autorization
}

// расширяем метод путём добавления информации о пользователе
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// get user from db
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	// Создаем токен методами библиотеки jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			// конец действия токена
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			// время создания токена
			IssuedAt: time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}
