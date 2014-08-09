package main

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/peterhellberg/generator-generator/server"
)

const (
	defaultPort = "1337"
)

func main() {
	// Open a connection to MongoDB
	log.Println("Connecting to MongoDB…")
	sess, err := mgo.Dial(mongoURL())
	if err != nil {
		log.Printf("Can't connect to MongoDB, go error %v\n", err)
		os.Exit(1)
	}
	defer sess.Close()

	port := getenv("PORT", defaultPort)
	err = server.ListenAndServe(os.Stdout, ":"+port)
	if err != nil {
		panic(err)
	}
}

func mongoURL() (url string) {
	url = os.Getenv("MONGOHQ_URL")

	if url == "" {
		log.Println("ENV variable MONGOHQ_URL not set!")
		os.Exit(1)
	}

	return
}

func getenv(key, fallback string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return fallback
}
