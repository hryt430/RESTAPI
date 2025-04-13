package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hryt430/RESTAPI/api/internal/domain/entity"
	"github.com/hryt430/RESTAPI/api/internal/interfaces/database"
	authService "github.com/hryt430/RESTAPI/api/internal/usecase/auth"
	userService "github.com/hryt430/RESTAPI/api/internal/usecase/user"
)

type AuthHandler struct {
	Interactor *authService.AuthDomainService
}

func NewAuthHandler(sqlHandler database.SqlHandler, authRepo authService.AuthServiceRepository) *AuthHandler {
	return &AuthHandler{
		Interactor: authService.NewAuthDomainService(
			authRepo,
			userService.UserDomainService{
				UserServiceRepository: &database.UserServiceRepository{
					SqlHandler: sqlHandler,
				},
			},
		),
	}
}

func (ah *AuthHandler) SignUp(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	createdUser, err := ah.Interactor.SignUp(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign up"})
		return
	}
	ctx.JSON(http.StatusOK, createdUser)
}

func (ah *AuthHandler) Login(ctx *gin.Context) {
	var request struct {
		ID       int    `json:"id"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := ah.Interactor.Login(request.ID, request.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (ah *AuthHandler) GenerateToken(ctx *gin.Context) {
	// ミドルウェアでセットしたユーザー情報を取得
	userInterface, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user := userInterface.(*entity.User) // 実体型にキャスト

	token, err := ah.Interactor.GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (ah *AuthHandler) Validate(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	user, err := ah.Interactor.Validate(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
