package server

import (
	"errors"
	"html"
	"net/http"
	"regexp"

	"github.com/peterhellberg/generate.name/generator"
)

var rp = regexp.MustCompile("^[a-z0-9-]+$")

func (s *Server) createHandler(r *http.Request, w http.ResponseWriter) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	slug := html.EscapeString(r.FormValue("slug"))
	name := html.EscapeString(r.FormValue("name"))

	if !rp.MatchString(slug) {
		return errors.New("invalid slug" + slug)
	}

	g := &generator.Generator{
		Slug: slug,
		Name: name,
	}

	if err := s.Create(g); err != nil {
		return err
	}

	http.Redirect(w, r, "/"+slug+"/edit", 301)

	return nil
}
