package server

import (
	"net/http"
	"strconv"

	"github.com/peterhellberg/generator-generator/generator"
)

func generateHandler(ctx *Context, r *http.Request, w http.ResponseWriter) error {
	slug, err := getSlug(r, "/generate")
	if err != nil {
		return err
	}

	q := r.URL.Query()

	n, err := strconv.Atoi(q.Get("n"))
	if err != nil || n > 100 {
		n = 5
	}

	sep := q.Get("sep")

	if sep == "br" {
		sep = "<br>"
	} else {
		sep = "\n"
	}

	sess := ctx.Session.Clone()
	defer sess.Close()

	c := sess.DB("").C("generators")

	g := generator.Generator{}
	err = c.FindId(slug).One(&g)
	if err != nil {
		return err
	}

	w.Write(g.GenerateNJoined(n, sep))

	return nil
}
