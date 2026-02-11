package main

import (
	"flag"
	"groupieee/api"
	"groupieee/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := flag.String("port", ":8080", "Port du serveur")
	flag.Parse()
	if err := api.LoadData(); err != nil {
		log.Fatal(err)
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() { <-sig; os.Exit(0) }()
	log.Fatal(server.Start(*port))
}
