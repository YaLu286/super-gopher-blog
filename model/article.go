package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type Article struct {
	ID       int64
	Title    string
	Text     string
	PostDate time.Time
}

var DB *sql.DB

func ConnectDB() error {
	var err error
	confStr := "host=localhost port=5432 user=postgres password=123 dbname=postgres sslmode=disable"
	DB, err = sql.Open("postgres", confStr)
	if err != nil {
		return err
	}
	return nil
}

func GetArticles(offset int, limit int) ([]Article, error) {

	queryStr := fmt.Sprintf("SELECT id, title, text, postdate FROM articles WHERE id BETWEEN %d AND %d;", offset, offset+limit)
	// fmt.Println(queryStr)
	rows, err := DB.Query(queryStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	resArticles := make([]Article, 0, limit)

	for rows.Next() {
		var a Article
		err = rows.Scan(&a.ID, &a.Title, &a.Text, &a.PostDate)
		if err != nil {
			return nil, err
		}
		resArticles = append(resArticles, a)
	}
	return resArticles, nil
}

func PostArticle(name string, text string) error {
	// queryStr := fmt.Sprintf("INSERT INTO articles VALUES ((SELECT max(id) FROM articles) + 1, '%s', '%s', now())", name, text)
	// fmt.Println(queryStr)
	// _, err := DB.Exec(queryStr)
	_, err := DB.Exec("INSERT INTO articles VALUES ((SELECT max(id) FROM articles) + 1, $1, $2, now())", name, text)
	if err != nil {
		return err
	}
	return nil
}
