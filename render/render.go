package render

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Variable struct {
	Name  string
	Value string
}

/*
Renders a page and outputs the new one with the selected internal variables replaced with
new values defined as {{ NameHere }} in the html file selected.

For example;

main.go
render.Page("test.html", []Variable{{Name: "test", Value: "New Text"}})

test.html
<p>{{ test }}</p>

Output
<p>New Text</p>
*/
func Page(w http.ResponseWriter, r *http.Request, File string, Replacements []Variable) {
	fmt.Println("Rendering " + File + " for " + r.RemoteAddr + " on " + time.Now().Format(time.UnixDate))

	bytes, err := os.ReadFile(File)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to locate the inputed template of " + File))
		return
	}
	page := string(bytes)

	for _, i := range Replacements {
		page = strings.ReplaceAll(page, "{{ "+i.Name+" }}", i.Value)
	}

	w.WriteHeader(200)
	w.Write([]byte(page))
}

/*
Renders / Sends a file to the requester
*/
func File(w http.ResponseWriter, r *http.Request, File string) {
	bytes, err := os.ReadFile(File)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to locate the inputed file of " + File))
		return
	}

	w.WriteHeader(200)
	w.Write(bytes)
}

/*
Sends html page to imedietly refresh the loaded current url with GET request
*/
func Refresh(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("<html><head><meta HTTP-EQUIV=\"refresh\" CONTENT=\"0\"></head></html>"))
}

func Url(data string) string {
	return url.QueryEscape(data)
}
