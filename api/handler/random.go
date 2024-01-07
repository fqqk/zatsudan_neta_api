package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"math/rand"
	"time"
)

// define Topic struct
type Topic struct {
	Topic string `json:"topic"`
}

var PURPOSE_ID := map[string]int{
	"all": 1,
	"p_c": 2,
	"b_r": 3,
}

// return random topic
func RandomHandler(w http.ResponseWriter, r *http.Request, purposeKey string) {
	// generate random number
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100) + 1

	// get random talk topic
	topic, err := getRandom(db, PURPOSE_ID[purposeKey], n)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// return json response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(topic)
}

func getRandom(db *sql.DB, purposeID int, n int) (Topic, error) {
	query := fmt.Sprintf("SELECT topic FROM topics WHERE purpose_id = %d AND id = %d;", purposeID, n)
	record := db.QueryRow(query)

	var topic string
	if err := record.Scan(&topic); err != nil {
		if err == sql.ErrNoRows {
			// エラーがないが結果がない場合の処理
			return Topic{}, fmt.Errorf("no topic found for purpose_id=%d and id=%d", purposeID, n)
		}
		// その他のエラーの場合
		return Topic{}, err
	}

	return Topic{Topic: topic}, nil
}