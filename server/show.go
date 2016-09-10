package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/peterhellberg/generate.name/generator"
)

var show = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/show.html",
))

// ShowGenerator contains the generator
// and if it is editable
type ShowGenerator struct {
	generator.Generator
	IsEditable bool
}

func (s *Server) showHandler(r *http.Request, w http.ResponseWriter) error {
	slug, err := getSlug(r, "")
	if err != nil {
		return err
	}

	if strings.HasSuffix(slug, ".json") {
		return s.showJSON(strings.TrimSuffix(slug, ".json"), r, w)
	}

	sess := s.Session.Clone()
	defer sess.Close()

	c := sess.DB("").C("generators")

	g := generator.Generator{}
	err = c.FindId(slug).One(&g)
	if err != nil {
		return err
	}

	g.SetGenFunc(s.newGenFunc(slug))

	keyParam := r.URL.Query().Get("key")
	editable := g.Key == "" || g.Key == keyParam || validBackdoorKey(keyParam)

	return show.Execute(w, ShowGenerator{g, editable})
}

func (s *Server) newGenFunc(slug string) func(string) string {
	return func(src string) string {
		gSlug := strings.TrimSuffix(strings.TrimPrefix(src, "[GENERATE "), "]")

		if gSlug == "" || gSlug == slug {
			return src
		}

		sess := s.Session.Clone()
		defer sess.Close()

		g := &generator.Generator{}

		err := sess.DB("").C("generators").FindId(gSlug).One(g)
		if err != nil {
			return src
		}

		return string(g.Generate())
	}
}

func (s *Server) showJSON(slug string, r *http.Request, w http.ResponseWriter) error {
	sess := s.Session.Clone()
	defer sess.Close()

	g := &generator.Generator{}

	err := sess.DB("").C("generators").FindId(slug).One(g)
	if err != nil {
		return err
	}

	g.SetGenFunc(s.newGenFunc(slug))

	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil || n < 1 || n > 100 {
		n = 10
	}

	response := ShowResponse{
		Meta: ShowMeta{
			Timestamp: time.Now(),
			Name:      g.Name,
		},
	}

	for _, b := range g.GenerateN(n) {
		response.Data = append(response.Data, string(b))
	}

	response.Meta.Count = len(response.Data)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}

type ShowResponse struct {
	Meta ShowMeta `json:"meta"`
	Data []string `json:"data"`
}

type ShowMeta struct {
	Timestamp time.Time `json:"timestamp"`
	Name      string    `json:"name"`
	Count     int       `json:"count"`
}
