package server

import (
	"net/http"
	"text/template"

	"github.com/peterhellberg/generator-generator/generator"
)

var edit = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/edit.html",
))

func editHandler(ctx *Context, r *http.Request, w http.ResponseWriter) error {
	slug, err := getSlug(r, "/edit")
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

	return edit.Execute(w, g)
}
