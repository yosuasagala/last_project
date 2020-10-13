package main

import (
	"fmt"
	"html/template"
	"lastproject/user"
	"log"
	"net/http"
	//"bitbucket.org/Hacktive8/user"
)

func main() {
	fmt.Println("hello")
	Template := template.Must(template.ParseGlob("files/var/templates/*"))
	user := user.New(Template)

	http.HandleFunc("/", user.HomeUser)
	http.HandleFunc("/homelogged", user.HomeUserLogged)
	http.HandleFunc("/about", user.AboutUser)
	http.HandleFunc("/aboutlogged", user.AboutUserLogged)
	http.HandleFunc("/contact", user.ContactUser)
	http.HandleFunc("/contactlogged", user.ContactUserLogged)
	http.HandleFunc("/articles", user.ArticlesUser)
	http.HandleFunc("/articles/add", user.AddArticlesUser)
	http.HandleFunc("/articles/edit", user.EditArticlesUser)
	http.HandleFunc("/articles/remove", user.RemoveArticlesUser)
	http.HandleFunc("/articles/publish", user.PublishArticlesUser)
	http.HandleFunc("/articles/unpublish", user.UnpublishArticlesUser)
	http.HandleFunc("/register", user.RegisterUser)
	http.HandleFunc("/login", user.LoginUser)
	http.HandleFunc("/logout", user.LogoutUser)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("error listern")
	}
}
