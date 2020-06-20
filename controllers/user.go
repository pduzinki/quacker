package controllers

import (
	"net/http"

	"quacker/views"
)

// UserC is a controller struct responsible for handling user resources
type UserC struct {
	HomepageView *views.View
	LoginView    *views.View
	SignupView   *views.View
}

// NewUserC creates new user controller
func NewUserC() *UserC {
	uc := UserC{
		HomepageView: views.NewView("views/user/homepage.gohtml"),
		LoginView:    views.NewView("views/user/login.gohtml"),
		SignupView:   views.NewView("views/user/signup.gohtml"),
	}

	return &uc
}

// GetHomepage handles GET /
func (uc *UserC) GetHomepage(w http.ResponseWriter, r *http.Request) {
	uc.HomepageView.Render(w, r)
}

// GetLogin handles GET /login
func (uc *UserC) GetLogin(w http.ResponseWriter, r *http.Request) {
	uc.LoginView.Render(w, r)
}

// PostLogin handles POST /login
func (uc *UserC) PostLogin(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// GetSignup handles GET /signup
func (uc *UserC) GetSignup(w http.ResponseWriter, r *http.Request) {
	uc.SignupView.Render(w, r)
}

// PostSignup handles POST /signup
func (uc *UserC) PostSignup(w http.ResponseWriter, r *http.Request) {
	// TODO
}
