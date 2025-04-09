package authService

import "github.com/hryt430/RESTAPI/api/internal/domain/entity"

type AuthServiceRepository interface {
	GenerateToken(params *entity.User) (string, error)
	Validate(token string) (*entity.User, error)
}
