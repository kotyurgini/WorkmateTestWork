package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/kotyurgini/WorkmateTestWork/internal/server"
	"github.com/kotyurgini/WorkmateTestWork/internal/storage"
)

func main() {
	storage := storage.NewRAMStorage()
	server := server.NewServer(storage, getPortFromEnv())
	server.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-c
	log.Printf("Signal received: %s", sig)
	server.Shutdown()
	storage.Close()
	os.Exit(0)
}

func getPortFromEnv() int {
	port := os.Getenv("APP_PORT")
	if port == "" {
		return 8080
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Invalid APP_PORT environment variable: %s", port)
	}
	return p
}
