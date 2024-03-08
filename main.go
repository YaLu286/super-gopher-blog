package main

import (
	"SuperGopherBlog/auth"
	"SuperGopherBlog/controller"
	"SuperGopherBlog/model"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/articles/create/", auth.Authorize(http.HandlerFunc(controller.CreateArticlePage))).Methods("GET")

	router.Handle("/articles/create/", auth.Authorize(http.HandlerFunc(controller.CreateArticleHandler))).Methods("POST")

	router.HandleFunc("/articles/{page:[0-9]+}", controller.ShowArticles)

	router.HandleFunc("/article/{id:[0-9]+}", controller.ShowArticle)

	router.HandleFunc("/login", controller.LoginPage).Methods("GET")

	router.HandleFunc("/login", controller.LoginHandler).Methods("POST")

	router.HandleFunc("/logout", controller.LogoutHandler).Methods("POST")

	router.HandleFunc("/register", controller.RegisterPage).Methods("GET")

	router.HandleFunc("/register", controller.RegisterHandler).Methods("POST")

	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", router)

	err := model.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	http.ListenAndServe(":8888", router)

}
