package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simple-snake/api"
	"simple-snake/snake"
	"syscall"
	"time"
)

func main() {
	s := snake.New()
	h := api.New(s)
	var port string
	flag.StringVar(&port, "port", "8080", "server port")
	flag.Parse()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: h,
	}

	// start server with a graceful shutdown that doesn't interrupt connections
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %s", err)
		}
	}()
	log.Printf("server started on: localhost%s", srv.Addr)

	<-done
	log.Print("server stopping")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %s", err)
	}
	log.Print("server stopped")
}
