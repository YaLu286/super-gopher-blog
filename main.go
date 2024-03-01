package main

import (
	// "SuperGopherBlog/auth"
	"SuperGopherBlog/controller"
	"SuperGopherBlog/model"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/articles/create/", controller.CreateArticle)

	router.HandleFunc("/articles", controller.ShowArticles)

	router.HandleFunc("/article/", controller.ShowArticle)

	router.HandleFunc("/login", controller.LoginHandler)

	err := model.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	http.ListenAndServe(":8888", router)

}
