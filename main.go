package main

import (
	"filebox/config"
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

	conf, err := config.Load("FileBox.conf")
	if err != nil {
		conf = config.New("FileBox.conf")
		conf.Add("Host", "127.0.0.1:80")
		conf.Add("StoragePath", "storage")
		conf.Save()
	}

	users, err := config.Load("Users.conf")
	if err != nil {
		users = config.New("Users.conf")
		users.Add("Admin", "12345")
		users.Save()
	}

	pages.Users = users

	pages.StoragePath = conf.Get("StoragePath")

	http.ListenAndServe(conf.Get("Host"), server)
}
