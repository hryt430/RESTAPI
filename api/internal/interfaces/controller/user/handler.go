package user

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hryt430/RESTAPI/api/internal/domain/entity"
	"github.com/hryt430/RESTAPI/api/internal/interfaces/database"
	userService "github.com/hryt430/RESTAPI/api/internal/usecase/user"
)

type UserHandler struct {
	Interactor userService.UserDomainService
}

func NewUserHandler(sqlHandler database.SqlHandler) *UserHandler {
	return &UserHandler{
		Interactor: userService.UserDomainService{
			UserServiceRepository: &database.UserServiceRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// GetUsers godoc
// @Summary ユーザー一覧を取得
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} entity.User "ユーザー一覧"
// @Router /auth/users [get]
func (handler *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := handler.Interactor.FindUser()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)

}

// GetUserById godoc
// @Summary ユーザーを取得
// @Tags user
// @Accept json
// @Produce json
// @Param request body entity.User true "ユーザー情報"
// @Success 201 {object} entity.User "作成されたユーザー情報"
// @Router /auth/users/{id} [get]
func (handler *UserHandler) GetUserById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := handler.Interactor.FindUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid user ID"})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

// CreateUser godoc
// @Summary 新規ユーザーを作成
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "ユーザーID"
// @Param request body entity.User true "ユーザー情報"
// @Success 200 {object} entity.User "編集成功"
// @Router /auth/users/ [post]
func (handler *UserHandler) CreateUser(ctx *gin.Context) {
	var requestUser *entity.User
	if err := ctx.ShouldBindJSON(&requestUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	requestUser.CreatedAt = now
	requestUser.UpdatedAt = now

	newUser, err := handler.Interactor.CreateUser(requestUser)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newUser)
}

// EditUser godoc
// @Summary ユーザー情報を編集
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "ユーザーID"
// @Success 200 {object} int "削除されたユーザーID"
// @Router /auth/users/{id} [put]
func (handler *UserHandler) EditUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}

	user_id, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	if id != user_id {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "permission denied"})
		return
	}

	var requestUser *entity.User
	if err := ctx.ShouldBindJSON(&requestUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	requestUser.UpdatedAt = now

	user, err := handler.Interactor.EditUser(id, requestUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary ユーザー情報を削除
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "ユーザーID"
// @Success 200 {object} entity.User "削除成功"
// @Router /auth/users/{id} [delete]
func (handler *UserHandler) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user_id, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	if id != user_id {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "permission denied"})
		return
	}

	err = handler.Interactor.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, id)
}
