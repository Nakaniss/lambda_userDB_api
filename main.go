package main

import (
	"lmd-func/db"
	"lmd-func/handlers"
	"lmd-func/repository"
	"lmd-func/routes"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// データベースの初期化
	db.Init()
	defer db.Close()

	// リポジトリの初期化
	userRepo := repository.NewUserRepository(db.DB)

	// ハンドラーの初期化
	userHandler := handlers.NewUserHandler(userRepo)

	// ルーターの初期化
	router := routes.NewRouter(userHandler)

	// Lambda 関数のハンドラーを登録
	lambda.Start(router.Route)
}
