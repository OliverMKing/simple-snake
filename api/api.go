package api

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-snake/snake"
)

type Api interface {
	Run()
}

type api struct {
	snake snake.Snake
	port  string
}

var _ Api = &api{}

func New(s snake.Snake, port string) Api {
	return &api{
		snake: s,
		port:  port,
	}
}

func (a *api) Run() {
	http.HandleFunc("/", a.infoHandler)
	log.Fatal(http.ListenAndServe(":"+a.port, nil))
}

func (a *api) infoHandler(w http.ResponseWriter, req *http.Request) {
	resp := a.snake.Info()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode((resp)); err != nil {
		log.Printf("failed to encode: %s", err)
	}
}
