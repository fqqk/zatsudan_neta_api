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
	env := os.Getenv("APP_ENV")

	var connStr string
	if env == "production" {
		// Heroku PostgresのDATABASE_URL環境変数を取得
		connStr = os.Getenv("DATABASE_URL")
	} else {
		dbUser := os.Getenv("DB_USER")
		dbName := os.Getenv("DB_NAME")
		dbPassword := os.Getenv("DB_PASSWORD")
	
		connStr = fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPassword)
	}

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

	fmt.Printf("Connected to the %s database\n", env)
}
