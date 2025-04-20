package presenter

import (
	"context"

	"github.com/gin-gonic/gin"
	jwt_auth "github.com/hryt430/RESTAPI/api/internal/infrastructure/auth"
	"github.com/hryt430/RESTAPI/api/internal/infrastructure/config"
	databaseInfra "github.com/hryt430/RESTAPI/api/internal/infrastructure/database"
	jwt "github.com/hryt430/RESTAPI/api/internal/infrastructure/middleware"
	"github.com/hryt430/RESTAPI/api/internal/interfaces/controller/auth"
	"github.com/hryt430/RESTAPI/api/internal/interfaces/controller/system"
	"github.com/hryt430/RESTAPI/api/internal/interfaces/controller/user"
)

// サーバー呼び出しのための構造体
type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(ctx context.Context) error {
	r := gin.Default()

	// 死活監視用
	{
		systemHandler := system.NewSystemHandler()
		r.GET("/health", systemHandler.Health)
	}

	key := config.JWT_SECRET
	jwtRepo := jwt_auth.NewJwtAuthRepository(key)

	dbi := databaseInfra.NewSqlHandler()
	defer dbi.Close()

	userController := user.NewUserHandler(dbi)
	authController := auth.NewAuthHandler(dbi, jwtRepo)

	r.POST("/signup", authController.SignUp)
	r.POST("/login", authController.Login)

	auth := r.Group("/auth")
	auth.Use(jwt.AuthMiddleware(jwtRepo))

	// ユーザー管理システム
	{
		auth.GET("/validate", authController.Validate)
		auth.GET("/users", func(ctx *gin.Context) { userController.GetUsers(ctx) })
		auth.GET("/users/:id", func(ctx *gin.Context) { userController.GetUserById(ctx) })
		auth.POST("/users", func(ctx *gin.Context) { userController.CreateUser(ctx) })
		auth.PUT("/users/:id", func(ctx *gin.Context) { userController.EditUser(ctx) })
		auth.DELETE("/users/:id", func(ctx *gin.Context) { userController.DeleteUser(ctx) })
	}

	// サーバー起動
	port := config.PORT
	err := r.Run(port)
	if err != nil {
		return err
	}

	return nil
}
