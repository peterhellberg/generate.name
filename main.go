package main

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/peterhellberg/generate.name/server"
)

const (
	defaultPort = "1337"
)

func main() {
	log.Println("Connecting to MongoDB…")
	sess, err := mgo.Dial(mongoURL())
	if err != nil {
		log.Printf("Can't connect to MongoDB, go error %v\n", err)
		os.Exit(1)
	}
	defer sess.Close()

	s := &server.Server{
		Logger:  log.New(os.Stdout, "", 0),
		Session: sess,
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
