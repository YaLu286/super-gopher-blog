package controller

import (
	"SuperGopherBlog/auth"
	"SuperGopherBlog/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type ViewData struct {
	Articles []model.Article
	PrevPage int
	NextPage int
	LastPage int
}

func ShowArticles(w http.ResponseWriter, r *http.Request) {
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

func CreateArticle(w http.ResponseWriter, r *http.Request) {

	if auth.IsAuthorized(w, r) {

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
			w.WriteHeader(http.StatusCreated)
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

	} else {
		// LoginHandler(w, r)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func ShowArticle(w http.ResponseWriter, r *http.Request) {
	articleID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || articleID < 0 || articleID > 1000 {
		http.NotFound(w, r)
		return
	}

	article, err := model.GetArticle(articleID)
	if err != nil {
		return
	}

	fmt.Println(article)

	tmpl, err := template.ParseFiles("./html/article.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tmpl.Execute(w, article)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		auth.Signin(w, r)
		// http.ServeFile(w, r, "html/success.html")
		// http.Redirect(w, r, "articles/create/", http.StatusPermanentRedirect)
		w.Write([]byte("OK"))
		return
	} else {
		http.ServeFile(w, r, "html/login.html")
	}

}
