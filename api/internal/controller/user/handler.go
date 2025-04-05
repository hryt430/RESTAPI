package user

import "github.com/gin-gonic/gin"

type UserHandler struct{}

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
// @Success 200 {object} []user.ResponseUser "ユーザー一覧"
// @Router /v1/users [get]
func (h *UserHandler) GetUsers(ctx *gin.Context) {}

// GetUserById godoc
// @Summary ユーザーの詳細情報を取得
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "ユーザーID"
// @Success 200 {object} user.ResponseUser "ユーザー詳細"
// @Router /v1/users/{id} [get]
func (h *UserHandler) GetUserById(ctx *gin.Context) {}

// EditUser godoc
// @Summary ユーザー情報を編集
// @Tags user
// @Accept json
// @Produce json
// @Param request body user.RequestUserParam true "ユーザー情報"
// @Success 200 {object} user.Response "編集成功"
// @Router /v1/users [post]
func (h *UserHandler) EditUser(ctx *gin.Context) {}

// DeleteUser godoc
// @Summary ユーザー情報を削除
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "ユーザーID"
// @Success 200 {object} user.Response "削除成功"
// @Router /v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(ctx *gin.Context) {}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
