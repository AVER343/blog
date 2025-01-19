package auth

import "github.com/golang-jwt/jwt/v5"

type JWTAutheticator struct {
	secret string
	aud    string
	iss    string
}

func NewJWTAutheticator(secret, aud, iss string) *JWTAutheticator {
	return &JWTAutheticator{
		secret, iss, aud,
	}
}

func (j *JWTAutheticator) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return nil, nil
	})
}

func (j *JWTAutheticator) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}
