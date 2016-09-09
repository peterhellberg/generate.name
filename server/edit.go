package server

import (
	"net/http"
	"text/template"

	"github.com/peterhellberg/generate.name/generator"
)

var edit = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/edit.html",
))

// EditGenerator contains the generator
// and if it is editable
type EditGenerator struct {
	generator.Generator
	IsEditable bool
}

func (s *Server) editHandler(r *http.Request, w http.ResponseWriter) error {
	slug, err := getSlug(r, "/edit")
	if err != nil {
		return err
	}

	sess := s.Session.Clone()
	defer sess.Close()

	c := sess.DB("").C("generators")

	g := generator.Generator{}
	err = c.FindId(slug).One(&g)
	if err != nil {
		return err
	}

	keyParam := r.URL.Query().Get("key")
	editable := g.Key == "" || g.Key == keyParam || validBackdoorKey(keyParam)

	return edit.Execute(w, EditGenerator{g, editable})
}
