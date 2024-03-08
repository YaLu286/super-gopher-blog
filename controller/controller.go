package controller

import (
	"SuperGopherBlog/auth"
	"SuperGopherBlog/model"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type ViewData struct {
	Articles []model.Article
	PrevPage int
	NextPage int
	LastPage int
}

func ShowArticles(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pageID, _ := strconv.Atoi(vars["page"])

	limit := 3
	offset := pageID * limit

	articles, err := model.GetArticles(offset, limit)
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl, err := template.ParseFiles("./templates/index.html", "./templates/styles.tmpl", "./templates/header.tmpl", "./templates/menu.tmpl", "./templates/footer.tmpl")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		fmt.Println(err)
		return
	}

	data := &ViewData{
		Articles: articles,
		PrevPage: pageID - 1,
		NextPage: pageID + 1,
		LastPage: 1000, // TO FIX!
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func CreateArticlePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/create.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	text := r.PostFormValue("text")

	err := model.PostArticle(title, text)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		fmt.Println(err)
		return
	}

	tmpl, err := template.ParseFiles("./templates/posted.html")
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

}

func ShowArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID, _ := strconv.Atoi(vars["id"])

	article, err := model.GetArticle(articleID)
	if err != nil {
		return
	}

	tmpl, err := template.ParseFiles("./templates/article.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tmpl.Execute(w, article)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/login.html", "./templates/styles.tmpl", "./templates/header.tmpl", "./templates/menu.tmpl", "./templates/footer.tmpl")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	var data struct {
		Message string
	}

	c, err := r.Cookie("login")
	if err == nil && c.Value == "retry" && c.Expires.Before(time.Now()) {
		data.Message = "Invalid login and(or) password. Please, try again..."
	}

	err = tmpl.Execute(w, &data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	res := auth.Signin(w, r)

	r.Method = "GET"

	if res {
		http.Redirect(w, r, "/articles/create/", http.StatusFound)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:    "login",
			Value:   "retry",
			Expires: time.Now().Add(time.Second * 10),
		})
		http.Redirect(w, r, "/login", http.StatusFound)
	}

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	auth.Logout(w, r)

	r.Method = "GET"

	http.Redirect(w, r, "/login", http.StatusFound)

}

func RegisterPage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./templates/register.html", "./templates/styles.tmpl", "./templates/header.tmpl", "./templates/menu.tmpl", "./templates/footer.tmpl")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	var data struct {
		Message string
	}

	c, err := r.Cookie("Register-Error")
	if err == nil && c.Value != "" && c.Expires.Before(time.Now()) {
		data.Message = c.Value
	}

	err = tmpl.Execute(w, &data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	err := auth.Register(w, r)

	if err != nil {

		fmt.Println(r.Method)

		r.Method = "GET"

		http.SetCookie(w, &http.Cookie{
			Name:    "Register-Error",
			Value:   err.Error(),
			Expires: time.Now().Add(time.Second * 10),
		})

		http.Redirect(w, r, "/register", http.StatusFound)
		return
	}

	http.ServeFile(w, r, "./templates/success.html")

}
