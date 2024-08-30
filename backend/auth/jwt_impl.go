package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type JwtServiceImpl struct {
}

func NewJwtService() JwtService {
	return &JwtServiceImpl{}
}

func (service JwtServiceImpl) GenerateToken(userID int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	SecretKey := []byte(os.Getenv("SECRET_KEY"))

	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (service JwtServiceImpl) ValidateToken(token string) (*jwt.Token, error) {
	parse, err := jwt.Parse(token, service.parse)
	if err != nil {
		return nil, err
	}
	return parse, nil
}
func (service JwtServiceImpl) parse(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return []byte(os.Getenv("SECRET_KEY")), nil
}
