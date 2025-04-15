package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_DSN     string
	JWT_SECRET string
	PORT       string
)

func init() { //main関数が呼ばれる前に呼ばれる
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .envファイルが見つかりませんでした。")
	}

	DB_DSN = os.Getenv("DB_DSN")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	PORT = os.Getenv("PORT")
}
