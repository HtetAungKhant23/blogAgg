package main

import (
	"net/http"
	"time"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, struct {
		Message string    `json:"message"`
		Time    time.Time `json:"time"`
	}{
		Message: "server is live",
		Time:    time.Now(),
	})
}
