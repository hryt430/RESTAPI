package authService

import (
	"github.com/hryt430/RESTAPI/api/internal/domain/entity"
	userService "github.com/hryt430/RESTAPI/api/internal/usecase/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthDomainService struct {
	AuthServiceRepository AuthServiceRepository
	UserDomainService     userService.UserDomainService
}

func NewAuthDomainService(AuthServiceRepository AuthServiceRepository, UserDomainService userService.UserDomainService) *AuthDomainService {
	return &AuthDomainService{AuthServiceRepository: AuthServiceRepository, UserDomainService: UserDomainService}
}

func (ads AuthDomainService) SignUp(user *entity.User) (*entity.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashed)

	createdUser, err := ads.UserDomainService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, err

}

func (ads AuthDomainService) Login(id int, password string) (string, error) {
	user, err := ads.UserDomainService.FindUserById(id)
	if err != nil || user == nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token, err := ads.AuthServiceRepository.GenerateToken(user)
	return token, err
}

func (ads AuthDomainService) GenerateToken(user *entity.User) (string, error) {
	return ads.AuthServiceRepository.GenerateToken(user)
}

func (ads AuthDomainService) Validate(token string) (int, error) {
	return ads.AuthServiceRepository.Validate(token)
}
