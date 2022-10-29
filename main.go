package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/peterhellberg/generate.name/server"
)

const (
	defaultPort = "1337"
)

func main() {
	log.Println("Connecting to MongoDB…")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL()))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	s := &server.Server{
		Logger: log.New(os.Stdout, "", 0),
		Client: client,
	}

	port := getenv("PORT", defaultPort)

	log.Printf("Starting serving on http://0.0.0.0:%s", port)
	if err := s.ListenAndServe(":" + port); err != nil {
		panic(err)
	}
}

func mongoURL() string {
	if url := os.Getenv("MONGO_URL"); url != "" {
		return url
	}

	return "mongodb://localhost/generator-generator"
}

func getenv(key, fallback string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return fallback
}
