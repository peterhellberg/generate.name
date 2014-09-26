package server

import (
	"errors"
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

	g := generator.Generator{}
	err = c.FindId(slug).One(&g)
	if err != nil {
		return err
	}

	keyParam := r.URL.Query().Get("key")
	editable := g.Key == "" || g.Key == keyParam

	if !editable {
		return errors.New("not allowed to edit this generator")
	}

	name := html.EscapeString(r.FormValue("name"))
	key := html.EscapeString(r.FormValue("key"))

	_, err = c.UpsertId(slug, &generator.Generator{
		Slug:     slug,
		Name:     name,
		Key:      key,
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

	path := "/" + slug + "/edit"

	if key != "" {
		path = path + "?key=" + key
	}

	http.Redirect(w, r, path, 301)

	return nil
}
