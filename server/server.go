package server

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"gopkg.in/mgo.v2"
)

var backdoorKey = os.Getenv("BACKDOOR_KEY")

func validBackdoorKey(keyParam string) bool {
	return backdoorKey != "" && keyParam == backdoorKey
}

// Context contains the logger and MongoDB session
type Context struct {
	Logger  *log.Logger
	Session *mgo.Session
}

// Handler takes a context, request and response writer
type Handler func(*Context, *http.Request, http.ResponseWriter) error

func handlerFunc(ctx *Context, fn Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := fn(ctx, r, w)
		if err != nil {
			ctx.Logger.Printf("handlerFunc: uri=%s err=%s", r.RequestURI, err)
			w.WriteHeader(500)
		}
	})
}

func getSlug(r *http.Request, suffix string) (string, error) {
	segments := strings.Split(strings.TrimSuffix(r.URL.Path[1:], suffix), "/")

	if len(segments) > 1 {
		return "", errors.New("not a valid path: " + r.URL.Path)
	}

	return segments[0], nil
}

// ListenAndServe creates a context, registers all handlers
// and starts listening on the provided addr
func ListenAndServe(ctx *Context, addr string) error {
	http.Handle("/", handlerFunc(ctx, routesHandler))

	http.HandleFunc("/favicon.png", favicon)
	http.HandleFunc("/favicon.ico", favicon)

	return http.ListenAndServe(addr, nil)
}
