package main

import (
	"log"
	"net/http"
	"./database.go"
	"./router.go"
)

func main() {
	// データベース接続
	InitDB()
	// ルーターの初期化
	router := InitRouter()

	// サーバーを起動
	log.Fatal(http.ListenAndServe(":8080", router))
}
