package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Password string
type User struct {
	ID        string   `json:"id"`
	Username  string   `json:"username"`
	Password  Password `json:"-"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Email     string   `json:"email"`
}

func NewUser(ID string, username, password, email string) *User {
	user := &User{
		Username:  username,
		Email:     email,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
		ID:        ID,
	}
	user.Password.Set(password)
	return user
}

func (p *Password) Set(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	*p = Password(hash)
	return nil
}

func (p *Password) ComparePassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
