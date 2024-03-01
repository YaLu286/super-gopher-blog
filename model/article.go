package model

import (
	"database/sql"
	// "fmt"
	"time"

	_ "github.com/lib/pq"
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

func GetArticle(ID int) (*Article, error) {
	row := DB.QueryRow("SELECT * FROM articles WHERE id = $1", ID)
	resArticle := &Article{}

	err := row.Scan(&resArticle.ID, &resArticle.Title, &resArticle.Text, &resArticle.PostDate)
	if err != nil {
		return nil, err
	}

	return resArticle, nil
}

func GetArticles(offset int, limit int) ([]Article, error) {
	rows, err := DB.Query(`SELECT * FROM articles 
							ORDER BY postdate DESC
							LIMIT $1 offset $2;`, limit, offset)
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
	_, err := DB.Exec("INSERT INTO articles VALUES ((SELECT max(id) FROM articles) + 1, $1, $2, now())", name, text)
	if err != nil {
		return err
	}
	return nil
}
