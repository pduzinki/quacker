package controllers

import (
	"github.com/gorilla/schema"
	"net/http"
)

type signupForm struct {
	Username string `schema:"username"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type loginForm struct {
	Login    string `schema:"login"`
	Password string `schema:"password"`
}

type quackForm struct {
	Text string `schema:"quacktext"`
}

func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(true)
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}

	return nil
}
