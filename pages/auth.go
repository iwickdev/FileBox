package pages

import (
	"filebox/auth"
	"filebox/config"
	"filebox/render"
	"net/http"
	"time"
)

var Users config.Configuration

func Auth(w http.ResponseWriter, r *http.Request) {
	if _, ok := auth.Validate(w, r); ok {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	r.ParseMultipartForm(1024)
	usrname := r.Form.Get("username")
	pswword := r.Form.Get("password")

	if (Users.Get(usrname) == "") || (pswword != Users.Get(usrname)) {
		w.WriteHeader(301)
		w.Write([]byte("Access Denied\nProvided credentials are invalid"))
		return
	}

	auth.Add(w, r, time.Now().Add(time.Minute*30))

	render.Refresh(w, r)
}
