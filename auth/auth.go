package auth

import (
	// "SuperGopherBlog/controller"
	"SuperGopherBlog/model"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var Sessions = map[string]Session{}

type Session struct {
	username string
	expiry   time.Time
}

func (s Session) IsExpired() bool {
	return s.expiry.Before(time.Now())
}

func Signin(w http.ResponseWriter, r *http.Request) {
	login := r.PostFormValue("login")
	password := r.PostFormValue("password")

	user, exists := model.FindUser(login)

	if !exists || CheckPasswordHash(password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("BAD CREDENTIALS!"))
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	Sessions[sessionToken] = Session{
		username: user.Login,
		expiry:   expiresAt,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_cookie",
		Value:   sessionToken,
		Expires: expiresAt,
	})

}

func IsAuthorized(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("session_cookie")
	if err != nil {
		if err == http.ErrNoCookie {
			return false

		}
		return false
	}

	sessionToken := c.Value

	fmt.Print(sessionToken)

	userSession, exist := Sessions[sessionToken]

	if !exist {
		return false
	}

	if userSession.IsExpired() {
		delete(Sessions, sessionToken)
		return false
	}

	return true
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
