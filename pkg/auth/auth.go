package auth

import "github.com/golang-jwt/jwt/v5"

type Authenticator interface {
	GenerateToken(claims jwt.Claims) (*jwt.Token, error)
	ValidateToken(claims jwt.Claims) (*jwt.Token, error)
}
