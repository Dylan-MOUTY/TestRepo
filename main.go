package main

import (
	"groupieee/api"
	"groupieee/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // pour le run en local
	}
	addr := ":" + port

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() { <-sig; os.Exit(0) }()

	go func() {
		if err := api.LoadData(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Server listening on", addr)
	log.Fatal(server.Start(addr))
}
