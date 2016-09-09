package server

import (
	"errors"
	"net/http"
	"strings"
)

func (s *Server) routesHandler(r *http.Request, w http.ResponseWriter) error {
	p := r.URL.Path

	if r.Method == "GET" {
		if p == "/" {
			return s.indexHandler(r, w)
		}

		if strings.HasSuffix(p, "/edit") && len(p) > 5 {
			return s.editHandler(r, w)
		}

		if strings.HasSuffix(p, "/generate") && len(p) > 9 {
			return s.generateHandler(r, w)
		}

		if strings.HasSuffix(p, "/") {
			http.Redirect(w, r, strings.TrimRight(p, "/"), 301)
		}

		return s.showHandler(r, w)
	}

	if r.Method == "POST" {
		if r.URL.Path == "/" {
			return s.createHandler(r, w)
		}

		return s.updateHandler(r, w)
	}

	return errors.New("no such handler")
}
