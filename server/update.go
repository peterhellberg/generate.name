package server

import (
	"html"
	"net/http"
	"strings"

	"github.com/peterhellberg/generator-generator/generator"
)

func updateHandler(ctx *Context, r *http.Request, w http.ResponseWriter) error {
	slug, err := getSlug(r, "")
	if err != nil {
		return err
	}

	err = r.ParseForm()
	if err != nil {
		return err
	}

	sess := ctx.Session.Clone()
	defer sess.Close()

	c := sess.DB("").C("generators")

	_, err = c.UpsertId(slug, &generator.Generator{
		Slug:     slug,
		Name:     html.EscapeString(r.FormValue("name")),
		Field1:   strings.Split(strings.TrimSpace(html.EscapeString(r.FormValue("field1"))), "\n"),
		Field2:   strings.Split(strings.TrimSpace(html.EscapeString(r.FormValue("field2"))), "\n"),
		Field3:   strings.Split(strings.TrimSpace(html.EscapeString(r.FormValue("field3"))), "\n"),
		Field4:   strings.Split(strings.TrimSpace(html.EscapeString(r.FormValue("field4"))), "\n"),
		Field5:   strings.Split(strings.TrimSpace(html.EscapeString(r.FormValue("field5"))), "\n"),
		Field6:   strings.Split(strings.TrimSpace(html.EscapeString(r.FormValue("field6"))), "\n"),
		Template: html.EscapeString(r.FormValue("template")),
	})

	if err != nil {
		return err
	}

	http.Redirect(w, r, "/"+slug+"/edit", 301)

	return nil
}
