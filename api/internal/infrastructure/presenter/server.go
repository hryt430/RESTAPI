package presenter

import (
	"context"

	"github.com/gin-gonic/gin"
	jwt_auth "github.com/hryt430/RESTAPI/api/internal/infrastructure/auth"
	"github.com/hryt430/RESTAPI/api/internal/infrastructure/db"
	jwt "github.com/hryt430/RESTAPI/api/internal/infrastructure/middleware"
	"github.com/hryt430/RESTAPI/api/internal/interfaces/controller/auth"
	"github.com/hryt430/RESTAPI/api/internal/interfaces/controller/system"
	"github.com/hryt430/RESTAPI/api/internal/interfaces/controller/user"
)

// サーバー呼び出しのための構造体
type Server struct{}

func (s *Server) Run(ctx context.Context) error {
	r := gin.Default()

	jwtRepo := jwt_auth.NewJwtAuthRepository("secret-key")
	systemController := system.NewSystemHandler()
	userController := user.NewUserHandler(db.NewSqlHandler())
	authController := auth.NewAuthHandler(db.NewSqlHandler(), jwtRepo)

	// サーバーの死活管理
	{
		r.GET("/health", systemController.Health)
	}

	r.POST("/signup", authController.SignUp)
	r.POST("login", authController.Login)

	auth := r.Group("/auth")
	auth.Use(jwt.AuthMiddleware(jwtRepo))

	// ユーザー管理システム
	{
		auth.GET("/validate", authController.Validate)
		auth.GET("/users", func(ctx *gin.Context) { userController.GetUsers(ctx) })
		auth.GET("/users/:id", func(ctx *gin.Context) { userController.GetUserById(ctx) })
		auth.POST("/users", func(ctx *gin.Context) { userController.CreateUser(ctx) })
		auth.POST("/users/:id", func(ctx *gin.Context) { userController.EditUser(ctx) })
		auth.DELETE("/users/:id", func(ctx *gin.Context) { userController.DeleteUser(ctx) })
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
