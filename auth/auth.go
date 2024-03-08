package auth

import (
	// "SuperGopherBlog/controller"
	"SuperGopherBlog/model"
	"errors"
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

func Register(w http.ResponseWriter, r *http.Request) error {
	login := r.PostFormValue("login")

	password := r.PostFormValue("password")

	passwordRepeat := r.PostFormValue("password_repeat")

	_, exists := model.FindUser(login)

	if exists {
		return errors.New("User already exists")
	}

	if password != passwordRepeat {
		return errors.New("Passwords doesn't match")
	}

	passHash, err := HashPassword(password)
	if err != nil {
		return err
	}

	u := model.User{Login: login, Password: passHash}

	err = u.SaveUser()
	if err != nil {
		return err
	}

	return nil
}

func Signin(w http.ResponseWriter, r *http.Request) bool {
	login := r.PostFormValue("login")
	password := r.PostFormValue("password")

	fmt.Println(HashPassword(password))
	user, exists := model.FindUser(login)

	if !exists || !CheckPasswordHash(password, user.Password) {
		return false
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
	return true
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_cookie")

	if err != nil {
		return
	}
	if _, exists := Sessions[c.Value]; exists {
		delete(Sessions, c.Value)
	}
}

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthorized(w, r) {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
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
