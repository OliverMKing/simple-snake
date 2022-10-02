package api

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-snake/model"
	"simple-snake/snake"
)

const (
	contentType     = "Content-Type"
	applicationJson = "application/json"
)

func New(s snake.Snake) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", infoHandler(s))
	mux.HandleFunc("/start", noOp)
	mux.HandleFunc("/move", moveHandler(s))
	mux.HandleFunc("/end", noOp)
	return mux
}

func infoHandler(s snake.Snake) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		resp := s.Info()
		w.Header().Set(contentType, applicationJson)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("failed to encode: %s", err)
			return
		}
	}
}

func moveHandler(s snake.Snake) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.GameReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("failed to decode: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp := s.Move(req)
		w.Header().Set(contentType, applicationJson)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("failed to encode: %s", err)
			return
		}
	}
}

func noOp(_ http.ResponseWriter, _ *http.Request) {}
