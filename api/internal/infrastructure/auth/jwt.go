package authInfra

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
func (repository *JwtAuthRepository) GenerateToken(user *entity.User) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": now.Add(time.Hour * 24).Unix(),
		"iat": now.Unix(), //発行時刻
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(repository.secretKey))
}

// JWTを検証する
func (repository *JwtAuthRepository) Validate(tokenString string) (int, error) {
	// 秘密鍵を知ることでjwt側が構文を解析してくれる
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 署名がHMAC方式であるかどうかを確かめる
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(repository.secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	// トークンが有効か（時間と署名の正しさ）
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { //
		sub, ok := claims["sub"].(float64)
		if !ok {
			return 0, fmt.Errorf("invalid subject in token")
		}
		return int(sub), nil
	}

	return 0, fmt.Errorf("invalid token")
}
