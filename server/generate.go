package server

import (
	"net/http"
	"strconv"
)

func (s *Server) generateHandler(r *http.Request, w http.ResponseWriter) error {
	slug, err := getSlug(r, "/generate")
	if err != nil {
		return err
	}

	q := r.URL.Query()

	n, err := strconv.Atoi(q.Get("n"))
	if err != nil || n > 15 {
		n = 5
	}

	sep := q.Get("sep")

	if sep == "br" {
		sep = "<br>"
	} else {
		sep = "\n"
	}

	g, err := s.Store.Find(slug)
	if err != nil {
		return err
	}

	g.SetGenFunc(s.newGenFunc(slug))

	_, err = w.Write(g.GenerateNJoined(n, sep))

	return err
}
