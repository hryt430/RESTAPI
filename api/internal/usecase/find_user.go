package usecase

import (
	"context"
	"errors"

	"github.com/hryt430/RESTAPI/api/internal/domain/administrator"
	"github.com/hryt430/RESTAPI/api/internal/domain/entity"
	"github.com/hryt430/RESTAPI/api/internal/domain/user"
)

var (
	INVALID_USER_ID = errors.New("invalid user id")
	INVALID_ADMIN   = errors.New("invalid admin")
)

type FindUserUsecase struct {
	uds user.UserDomainService
	ads administrator.AdminDomainService
}

func NewFindUserUseCase(uds user.UserDomainService, ads administrator.AdminDomainService) *FindUserUsecase {
	return &FindUserUsecase{uds: uds, ads: ads}
}

func (us *FindUserUsecase) Run(ctx context.Context) ([]*entity.User, error) {
	value, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, INVALID_USER_ID
	}

	AdminUser, err := us.ads.FindUser(ctx, value)
	if err != nil {
		return nil, err
	}

	if AdminUser.GetAdmin() != 1 {
		return nil, INVALID_ADMIN
	}

	users, err := us.uds.FindUser(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
