// Package cookies provides functions to manage HTTP cookies.
package cookies

import (
	"fmt"
	"net/http"
	"strconv"
)

type Status struct {
	Code    int
	Message string
}

type cookieName string

const (
	cookieNameStatusMessage cookieName = "status-message"
	cookieNameStatusCode    cookieName = "status-code"
)

func SetStatus(w http.ResponseWriter, code int, msg string) {
	set(w, string(cookieNameStatusCode), fmt.Sprint(code), 0)
	set(w, string(cookieNameStatusMessage), msg, 0)
}

func GetStatus(r *http.Request) (*Status, error) {
	codeCookie, err := r.Cookie(string(cookieNameStatusCode))
	if err != nil {
		return nil, err
	}
	messageCookie, err := r.Cookie(string(cookieNameStatusMessage))
	if err != nil {
		return nil, err
	}
	code, err := strconv.Atoi(codeCookie.Value)
	if err != nil {
		return nil, err
	}
	return &Status{Code: code, Message: messageCookie.Value}, nil
}

func ClearStatus(w http.ResponseWriter) {
	set(w, string(cookieNameStatusCode), "", -1)
	set(w, string(cookieNameStatusMessage), "", -1)
}

func set(w http.ResponseWriter, name string, value string, maxAge int) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   maxAge,
		HttpOnly: true,
	})
}
