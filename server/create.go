package server

import "net/http"

func (s *Server) createHandler(r *http.Request, w http.ResponseWriter) error {
	// err := r.ParseForm()
	// if err != nil {
	// 	return err
	// }

	// slug := html.EscapeString(r.FormValue("slug"))

	// rp := regexp.MustCompile("^[a-z0-9-]+$")

	// if !rp.MatchString(slug) {
	// 	return errors.New("invalid slug" + slug)
	// }

	// sess := s.Session.Clone()
	// defer sess.Close()

	// c := sess.DB("").C("generators")

	// _, err = c.UpsertId(slug, bson.M{"$set": bson.M{
	// 	"name": html.EscapeString(r.FormValue("name")),
	// }})

	// if err != nil {
	// 	return err
	// }

	// http.Redirect(w, r, "/"+slug+"/edit", 301)

	return nil
}
