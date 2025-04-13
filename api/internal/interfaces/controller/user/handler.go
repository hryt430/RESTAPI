package user

import (
	"errors"
	"strconv"

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

type Response struct {
	Status  string `json:"status"`  // ステータス（成功、失敗など）
	Message string `json:"message"` // メッセージ（詳細説明）
}

type ResponseUser struct {
	ID    string `json:"id"`    // ユーザーID
	Name  string `json:"name"`  // ユーザー名
	Email string `json:"email"` // メールアドレス
}

type RequestUserParam struct {
	Name  string `json:"name" binding:"required"`  // ユーザー名（必須）
	Email string `json:"email" binding:"required"` // メールアドレス（必須）
}

// GetUsers godoc
// @Summary ユーザー一覧を取得
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} []ResponseUser "ユーザー一覧"
// @Router /v1/users [get]
func (handler *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := handler.Interactor.FindUser()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, users)

}

// GetUserById godoc
// @Summary ユーザーの詳細情報を取得
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "ユーザーID"
// @Success 200 {object} ResponseUser "ユーザー詳細"
// @Router /v1/users/{id} [get]
func (handler *UserHandler) GetUserById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := handler.Interactor.FindUserById(id)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, user)

}

// EditUser godoc
// @Summary ユーザー情報を編集
// @Tags user
// @Accept json
// @Produce json
// @Param request body RequestUserParam true "ユーザー情報"
// @Success 200 {object} Response "編集成功"
// @Router /v1/users [post]
func (handler *UserHandler) EditUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	var requestUser *entity.User
	if err := ctx.ShouldBindJSON(&requestUser); err != nil {
		ctx.JSON(400, err)
		return
	}

	user, err := handler.Interactor.EditUser(id, requestUser)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, user)
}

// DeleteUser godoc
// @Summary ユーザー情報を削除
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "ユーザーID"
// @Success 200 {object} user.Response "削除成功"
// @Router /v1/users/{id} [delete]
func (handler *UserHandler) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, errors.New("無効なIDフォーマットです"))
		return
	}

	err = handler.Interactor.DeleteUser(id)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, id)
}

// CreateUser godoc
// @Summary 新規ユーザーを作成
// @Tags user
// @Accept json
// @Produce json
// @Param request body user.RequestUserParam true "ユーザー情報"
// @Success 201 {object} user.ResponseUser "作成されたユーザー情報"
func (handler *UserHandler) CreateUser(ctx *gin.Context) {
	var requestUser *entity.User
	if err := ctx.ShouldBindJSON(&requestUser); err != nil {
		ctx.JSON(400, err)
		return
	}

	newUser, err := handler.Interactor.CreateUser(requestUser)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, newUser)
}
