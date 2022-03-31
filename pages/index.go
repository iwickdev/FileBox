package pages

import (
	"filebox/auth"
	"filebox/render"
	"fmt"
	"io"
	"net/http"
	"os"
)

var StoragePath string

func Index(w http.ResponseWriter, r *http.Request) {
	sess, loggedin := auth.Validate(w, r)
	if !loggedin {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	if r.Method == "POST" {
		file, header, err := r.FormFile("upfile")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unable to read file form data"))
			return
		}

		dest, err := os.Create(StoragePath + r.URL.Path + "/" + header.Filename)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unable to read file stream"))
			return
		}

		io.Copy(dest, file)
		file.Close()
		dest.Close()

		render.Refresh(w, r)
		return
	}

	if r.URL.Query().Get("logout") == "true" {
		auth.Invalidate(sess)
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	if r.URL.Query().Get("rmf") != "" {
		err := os.Remove(StoragePath + r.URL.Path + "/" + r.URL.Query().Get("rmf"))
		if err != nil {
			os.RemoveAll(StoragePath + r.URL.Path + "/" + r.URL.Query().Get("rmf"))
		}
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	if r.URL.Query().Get("dnl") != "" {
		bytes, err := os.ReadFile(StoragePath + r.URL.Path + "/" + r.URL.Query().Get("dnl"))
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("The file requested could not be found on disk"))
			return
		}
		w.WriteHeader(200)
		w.Write(bytes)
		return
	}

	if r.URL.Query().Get("md") != "" {
		err := os.Mkdir(StoragePath+r.URL.Path+"/"+r.URL.Query().Get("md"), os.ModePerm)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to create directory on server"))
		}
		http.Redirect(w, r, r.URL.Path, http.StatusTemporaryRedirect)
		return
	}

	fi, err := os.ReadDir(StoragePath + r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to read storage directory"))
		return
	}

	var files string
	for _, f := range fi {
		fsinfo, _ := f.Info()

		if fsinfo.IsDir() {
			files += "<tr>"
			files += " <td>" + f.Name() + "</td>"
			files += " <td>Directory</td>"
			files += " <td>"
			files += "  <a href=\"" + r.URL.Path + f.Name() + "\"><button class=\"options-btn\">View</button></a>"
			files += "  <a href=\"?rmf=" + render.Url(f.Name()) + "\"><button class=\"options-btn\">Delete</button></a>"
			files += " </td>"
			files += "</tr>"
		}
	}

	for _, f := range fi {
		fsinfo, _ := f.Info()

		if !fsinfo.IsDir() {
			files += "<tr>"
			files += " <td>" + f.Name() + "</td>"
			files += " <td>" + fmt.Sprint(float64(fsinfo.Size())*0.001) + " KB</td>"
			files += " <td>"
			files += "  <a href=\"?dnl=" + render.Url(f.Name()) + "\"><button class=\"options-btn\">View</button></a>"
			files += "  <a href=\"?rmf=" + render.Url(f.Name()) + "\"><button class=\"options-btn\">Delete</button></a>"
			files += " </td>"
			files += "</tr>"
		}
	}

	render.Page(w, r, "accets/index.html", []render.Variable{
		{
			Name:  "fileslist",
			Value: files,
		},
	})
}
