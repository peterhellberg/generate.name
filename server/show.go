package server

import (
	"net/http"
	"text/template"

	"github.com/peterhellberg/generator-generator/generator"
)

var show = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/show.html",
))

type ShowGenerator struct {
	generator.Generator
	IsEditable bool
}

func showHandler(ctx *Context, r *http.Request, w http.ResponseWriter) error {
	slug, err := getSlug(r, "")
	if err != nil {
		return err
	}

	sess := ctx.Session.Clone()
	defer sess.Close()

	c := sess.DB("").C("generators")

	g := generator.Generator{}
	err = c.FindId(slug).One(&g)
	if err != nil {
		return err
	}

	keyParam := r.URL.Query().Get("key")
	editable := g.Key == "" || g.Key == keyParam || validBackdoorKey(keyParam)

	return show.Execute(w, ShowGenerator{g, editable})
}
