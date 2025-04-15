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

// SignUp godoc
// @Summary ユーザー登録
// @Tags auth
// @Accept json
// @Produce json
// @Param request body entity.User true "ユーザー情報"
// @Success 200 {object} entity.User "作成されたユーザー情報"
// @Router /signup [post]
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

// Login godoc
// @Summary ログイン
// @Tags auth
// @Accept json
// @Produce json
// @Param request body object true "ログイン情報（IDとパスワード）"
// @Success 200 {object} map[string]string "アクセストークン"
// @Router /login [post]
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
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials" + token})
		return
	}
	ctx.Header("Authorization", "Bearer "+token)
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// GenerateToken godoc
// @Summary トークンを再生成
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "再生成されたトークン"
func (ah *AuthHandler) GenerateToken(ctx *gin.Context) {
	// ミドルウェアでセットしたユーザー情報を取得
	userInterface, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// ユーザー型にキャスト
	user := userInterface.(*entity.User)

	token, err := ah.Interactor.GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// Validate godoc
// @Summary トークン検証
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer トークン"
// @Success 200 {object} entity.User "認証済みユーザー情報"
// @Router /auth/validate [get]
func (ah *AuthHandler) Validate(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	id, err := ah.Interactor.Validate(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	user, err := ah.Interactor.UserDomainService.FindUserById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "falid to find user"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
