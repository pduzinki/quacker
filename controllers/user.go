package controllers

import (
	"net/http"

	"quacker/views"
)

type UserC struct {
	LoginView  *views.View
	SignupView *views.View
}

func NewUserC() *UserC {
	uc := UserC{
		LoginView:  views.NewView("views/user/login.gohtml"),
		SignupView: views.NewView("views/user/signup.gohtml"),
	}

	return &uc
}

// Login handles GET /login
func (uc *UserC) GetLogin(w http.ResponseWriter, r *http.Request) {
	uc.LoginView.Render(w, r)
}

// LoginPost handles POST /login
func (uc *UserC) PostLogin(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// Signup handles GET /signup
func (uc *UserC) GetSignup(w http.ResponseWriter, r *http.Request) {
	uc.SignupView.Render(w, r)
}

// SignupPost handles POST /signup
func (uc *UserC) PostSignup(w http.ResponseWriter, r *http.Request) {
	// TODO
}
