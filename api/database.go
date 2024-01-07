package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

// db データベース接続用変数
var db *sql.DB

// init 初期化関数でデータベースに接続
func initDB() {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	connStr := "user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// データベース接続確認
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
