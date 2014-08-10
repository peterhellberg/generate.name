package server

import (
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
		Name:     r.FormValue("name"),
		Field1:   strings.Split(r.FormValue("field1"), "\r"),
		Field2:   strings.Split(r.FormValue("field2"), "\r"),
		Field3:   strings.Split(r.FormValue("field3"), "\r"),
		Field4:   strings.Split(r.FormValue("field4"), "\r"),
		Field5:   strings.Split(r.FormValue("field5"), "\r"),
		Field6:   strings.Split(r.FormValue("field6"), "\r"),
		Template: r.FormValue("template"),
	})

	if err != nil {
		return err
	}

	http.Redirect(w, r, "/"+slug+"/edit", 301)

	return nil
}
