package auth

import "github.com/dgrijalva/jwt-go"

type JwtService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
