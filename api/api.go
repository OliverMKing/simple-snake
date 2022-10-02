package api

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-snake/snake"
)

func New(s snake.Snake) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", infoHandler(s))
	return mux
}

func infoHandler(s snake.Snake) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		resp := s.Info()
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode((resp)); err != nil {
			log.Printf("failed to encode: %s", err)
		}
	}
}
