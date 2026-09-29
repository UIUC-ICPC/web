package handlers

import (
	"encoding/json"
	"net/http"
	"sigicpc-backend/internal/db"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.Posts)
}
