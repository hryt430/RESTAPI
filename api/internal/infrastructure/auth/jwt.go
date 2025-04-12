package jwt_auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hryt430/RESTAPI/api/internal/domain/entity"
)

type JwtAuthRepository struct {
	secretKey string
}

func NewJwtAuthRepository(secretKey string) *JwtAuthRepository {
	return &JwtAuthRepository{secretKey: secretKey}
}

// JWTを生成する
func (repo *JwtAuthRepository) GenerateToken(user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(repo.secretKey))
}

// JWTを検証する
func (repo *JwtAuthRepository) Validate(tokenString string) (*entity.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method should be expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(repo.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := &entity.User{
			ID: int(claims["sub"].(float64)),
		}
		return user, nil
	}

	return nil, fmt.Errorf("invalid token")
}
