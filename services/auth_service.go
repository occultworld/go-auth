package services

import (
	"errors"
	"go-auth/models"
	"go-auth/utils"

	"golang.org/x/crypto/bcrypt"

	"github.com/jmoiron/sqlx"
)

type AuthService struct {
	DB *sqlx.DB
}

func (s *AuthService) Register(user models.User) error {
	var existingUser models.User
	err := s.DB.Get(&existingUser, "SELECT * FROM users WHERE username=$1", user.Username)
	if err == nil {
		return errors.New("username already taken")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	_, err = s.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, hashedPassword)
	return err
}

func (s *AuthService) Login(username, password string) (string, error) {
	var user models.User
	err := s.DB.Get(&user, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, _ := utils.GenerateJWT(username)
	return token, nil
}