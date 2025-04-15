package main

import (
	"context"

	"github.com/hryt430/RESTAPI/api/internal/infrastructure/presenter"
)

// @title ユーザー管理API
// @version 1.0
// @description ユーザー管理サーバーの起動
// @host localhost:8080
// @BasePath /
func main() {
	server := presenter.NewServer()
	if err := server.Run(context.Background()); err != nil {
		panic(err)
	}
}
