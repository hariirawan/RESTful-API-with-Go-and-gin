package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct{}

var SECRET_KEY = []byte("980bdfc68476d963c55b22d1a16d41f32858a936b8fb34c08bba739f392b4b0b")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
