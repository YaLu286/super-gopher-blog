package main

import (
	"SuperGopherBlog/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type ViewData struct {
	Articles []model.Article
	PrevPage int
	NextPage int
	LastPage int
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/articles/create/", createArticle)

	router.HandleFunc("/articles", showArticles)

	err := model.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	http.ListenAndServe(":8888", router)
}

func showArticles(w http.ResponseWriter, r *http.Request) {
	pageID, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || pageID < 0 || pageID > 1000 {
		http.NotFound(w, r)
		return
	}

	limit := 3
	offset := pageID * limit

	articles, err := model.GetArticles(offset, limit)
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl, err := template.ParseFiles("./html/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		fmt.Println(err)
		return
	}

	data := &ViewData{
		Articles: articles,
		PrevPage: pageID - 1,
		NextPage: pageID + 1,
		LastPage: 1000,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		title := r.PostFormValue("title")
		text := r.PostFormValue("text")
		err := model.PostArticle(title, text)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			fmt.Println(err)
			return
		}
		tmpl, err := template.ParseFiles("./html/posted.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			fmt.Println(err)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal Server Error", 500)
		}

	} else {
		tmpl, err := template.ParseFiles("./html/create.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			fmt.Println(err)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal Server Error", 500)
		}
	}

}
