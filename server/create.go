package server

import (
	"errors"
	"html"
	"net/http"
	"regexp"

	"gopkg.in/mgo.v2/bson"
)

func createHandler(ctx *Context, r *http.Request, w http.ResponseWriter) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	slug := html.EscapeString(r.FormValue("slug"))

	rp := regexp.MustCompile("^[a-z0-9-]+$")

	if !rp.MatchString(slug) {
		return errors.New("invalid slug" + slug)
	}

	sess := ctx.Session.Clone()
	defer sess.Close()

	c := sess.DB("").C("generators")

	_, err = c.UpsertId(slug, bson.M{"$set": bson.M{
		"name": html.EscapeString(r.FormValue("name")),
	}})

	if err != nil {
		return err
	}

	http.Redirect(w, r, "/"+slug+"/edit", 301)

	return nil
}
