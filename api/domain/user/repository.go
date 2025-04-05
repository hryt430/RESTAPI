package user

import (
	"context"

	"github.com/hryt430/RESTAPI/api/domain/entity"
)

type UserServiceRepository interface {
	FindUser(c context.Context) ([]*entity.User, error)
	FindUserById(c context.Context, id string) (*entity.User, error)
	Save(c context.Context, params *entity.User) error
	Delete(c context.Context, id string) error
}
