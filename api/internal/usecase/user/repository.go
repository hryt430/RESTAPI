package userService

import (
	"github.com/hryt430/RESTAPI/api/internal/domain/entity"
)

type UserServiceRepository interface {
	FindUser() ([]*entity.User, error)
	FindUserById(id int) (*entity.User, error)
	Save(params *entity.User) (int, error)
	Edit(id int, params *entity.User) (int, error)
	Delete(id int) error
}
