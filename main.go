package main

import (
	"filebox/pages"
	"net/http"
	"os"
)

func main() {
	server := http.NewServeMux()

	os.Mkdir(pages.StoragePath, os.ModePerm)

	server.HandleFunc("/accets/", pages.Accets)
	server.HandleFunc("/login", pages.Login)
	server.HandleFunc("/auth", pages.Auth)
	server.HandleFunc("/", pages.Index)

	http.ListenAndServe("127.0.0.1:80", server)
}
