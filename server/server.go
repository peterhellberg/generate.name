package server

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/peterhellberg/generate.name/generator"
)

var backdoorKey = os.Getenv("BACKDOOR_KEY")

func validBackdoorKey(keyParam string) bool {
	return backdoorKey != "" && keyParam == backdoorKey
}

// Server contains the logger and MongoDB session
type Server struct {
	Logger *log.Logger
	Client *mongo.Client
}

// Handler takes a context, request and response writer
type Handler func(*http.Request, http.ResponseWriter) error

func (s *Server) handlerFunc(fn Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := fn(r, w)
		if err != nil {
			s.Logger.Printf("handlerFunc: uri=%s err=%s", r.RequestURI, err)
			w.WriteHeader(500)
		}
	})
}

// ListenAndServe creates a context, registers all handlers
// and starts listening on the provided addr
func (s *Server) ListenAndServe(addr string) error {
	http.Handle("/", s.handlerFunc(s.routesHandler))

	http.HandleFunc("/script.js", script)
	http.HandleFunc("/favicon.png", favicon)
	http.HandleFunc("/favicon.ico", favicon)

	return http.ListenAndServe(addr, nil)
}

func (s *Server) newGenFunc(slug string) func(string) string {
	return func(src string) string {
		gSlug := strings.TrimSuffix(strings.TrimPrefix(src, "[GENERATE "), "]")

		if gSlug == "" || gSlug == slug {
			return src
		}

		// sess := s.Session.Clone()
		// defer sess.Close()

		g := &generator.Generator{}

		// err := sess.DB("").C("generators").FindId(gSlug).One(g)
		// if err != nil {
		// 	return src
		// }

		return string(g.Generate())
	}
}

func getSlug(r *http.Request, suffix string) (string, error) {
	segments := strings.Split(strings.TrimSuffix(r.URL.Path[1:], suffix), "/")

	if len(segments) > 1 {
		return "", errors.New("not a valid path: " + r.URL.Path)
	}

	return segments[0], nil
}
