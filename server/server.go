package server

import (
	"io"
	"log"
	"net/http"
)

type Context struct {
	Logger *log.Logger
}

type Handler func(*Context, *http.Request, http.ResponseWriter) error

func baseHandler(ctx *Context, fn Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := fn(ctx, r, w)
		if err != nil {
			ctx.Logger.Printf("baseHandler: uri=%s err=%s", r.RequestURI, err)
			w.WriteHeader(500)
		}
	})
}

func helloHandler(ctx *Context, r *http.Request, w http.ResponseWriter) error {
	ctx.Logger.Printf("Obey the cock!")
	w.Write([]byte("Hello"))

	return nil
}

// ListenAndServe creates a context, registers all handlers
// and starts listening on the provided addr
func ListenAndServe(w io.Writer, addr string) error {
	ctx := &Context{
		Logger: log.New(w, "", 0),
	}

	http.Handle("/hello/", baseHandler(ctx, helloHandler))
	http.Handle("/", http.FileServer(http.Dir("./public")))

	return http.ListenAndServe(addr, nil)
}
