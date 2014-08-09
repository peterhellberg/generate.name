package main

import (
	"os"

	"github.com/peterhellberg/gen/server"
)

const (
	defaultPort = "1337"
)

func main() {
	port := getenv("PORT", defaultPort)
	err := server.ListenAndServe(os.Stdout, ":"+port)
	if err != nil {
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
