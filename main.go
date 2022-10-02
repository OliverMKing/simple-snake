package main

import (
	"simple-snake/api"
	"simple-snake/snake"
)

func main() {
	s := snake.New()
	server := api.New(s, "8080")
	server.Run()
}
