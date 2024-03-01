package model

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type User struct {
	Login    string
	Password string
}

func FindUser(login string) (*User, bool) {
	row := DB.QueryRow("SELECT * FROM users WHERE login=$1", login)
	u := &User{}
	err := row.Scan(&u.Login, &u.Password)
	if err == sql.ErrNoRows {
		return nil, false
	}
	return u, true
}
