package presenter

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hryt430/RESTAPI/api/internal/controller/system"
	"github.com/hryt430/RESTAPI/api/internal/controller/user"
)

// バージョン管理
const latest = "/v1"

// サーバー呼び出しのための構造体
type Server struct{}

func (s *Server) Run(c context.Context) error {
	r := gin.Default()
	v1 := r.Group(latest)

	// サーバーの死活管理
	{
		systemHandler := system.NewSystemHandler()
		v1.GET("/health", systemHandler.Health)
	}

	// ユーザー管理システム
	{
		userHandler := user.NewUserHandler()
		v1.GET("", userHandler.GetUsers)
		v1.GET("/:id", userHandler.GetUserById)
		v1.POST("", userHandler.EditUser)
		v1.DELETE("/:id", userHandler.DeleteUser)
	}

	// サーバー起動
	err := r.Run()
	if err != nil {
		return err
	}

	return nil
}

func NewServer() *Server {
	return &Server{}
}
