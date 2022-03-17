package pages

import (
	"filebox/auth"
	"filebox/render"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if _, ok := auth.Validate(w, r); ok {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	render.Page(w, r, "accets/login.html", []render.Variable{})
}
