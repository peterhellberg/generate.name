package main

import (
	"log"
	"os"

	"github.com/peterhellberg/generate.name/server"
	"github.com/peterhellberg/generate.name/stores/memory"
)

const defaultPort = "1337"

func main() {
	store := memory.NewStore()

	s := &server.Server{
		Logger: log.New(os.Stdout, "", 0),
		Store:  store,
	}

	port := getenv("PORT", defaultPort)

	log.Printf("Starting serving on http://0.0.0.0:%s", port)
	if err := s.ListenAndServe(":" + port); err != nil {
		panic(err)
	}
}

func getenv(key, fallback string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return fallback
}
