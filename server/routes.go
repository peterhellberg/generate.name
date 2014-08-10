package server

import (
	"errors"
	"net/http"
	"strings"
)

func routesHandler(ctx *Context, r *http.Request, w http.ResponseWriter) error {
	p := r.URL.Path

	if r.Method == "GET" {
		if p == "/" {
			return indexHandler(ctx, r, w)
		}

		if strings.HasSuffix(p, "/edit") && len(p) > 5 {
			return editHandler(ctx, r, w)
		}

		if strings.HasSuffix(p, "/generate") && len(p) > 9 {
			return generateHandler(ctx, r, w)
		}

		if strings.HasSuffix(p, "/") {
			http.Redirect(w, r, strings.TrimRight(p, "/"), 301)
		}

		return showHandler(ctx, r, w)
	}

	if r.Method == "POST" {
		if r.URL.Path == "/" {
			return createHandler(ctx, r, w)
		}

		return updateHandler(ctx, r, w)
	}

	return errors.New("no such handler")
}
