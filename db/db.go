package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQLドライバをインポート
)

var DB *sql.DB

// Init は、データベースへの接続を初期化します
func Init() {
	// 環境変数からデータベース接続情報を取得
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// データベースに接続
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)
	var err error
	DB, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 接続をプールする
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)

	// 接続が正常に行われたことを確認
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Connected to database")
}

// Close は、データベースへの接続を閉じます
func Close() {
	if DB != nil {
		DB.Close()
	}
}
