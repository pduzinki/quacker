package controllers

import (
	// "fmt"
	"net/http"

	"quacker/models"
	"quacker/views"
)

// UserController is responsible for handling user resources
type UserController struct {
	HomepageView *views.View
	LoginView    *views.View
	SignupView   *views.View
	us           models.UserService
}

// NewUserController creates user controller instance
func NewUserController(us models.UserService) *UserController {
	uc := UserController{
		HomepageView: views.NewView("views/user/homepage.gohtml"),
		LoginView:    views.NewView("views/user/login.gohtml"),
		SignupView:   views.NewView("views/user/signup.gohtml"),
		us:           us,
	}

	return &uc
}

// GetHomepage handles GET /
func (uc *UserController) GetHomepage(w http.ResponseWriter, r *http.Request) {
	uc.HomepageView.Render(w, r, nil)
}

// GetLogin handles GET /login
func (uc *UserController) GetLogin(w http.ResponseWriter, r *http.Request) {
	uc.LoginView.Render(w, r, nil)
}

// PostLogin handles POST /login
func (uc *UserController) PostLogin(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// GetSignup handles GET /signup
func (uc *UserController) GetSignup(w http.ResponseWriter, r *http.Request) {
	uc.SignupView.Render(w, r, nil)
}

// PostSignup handles POST /signup
func (uc *UserController) PostSignup(w http.ResponseWriter, r *http.Request) {
	var form signupForm
	var d views.Data

	err := parseForm(r, &form)
	if err != nil {
		d.SetAlert(err)
		uc.SignupView.Render(w, r, d)
		return
	}

	user := models.User{
		Username: form.Username,
		Email:    form.Email,
		Password: form.Password,
	}

	err = uc.us.Create(&user)
	if err != nil {
		d.SetAlert(err)
		uc.SignupView.Render(w, r, d)
	}

	http.Redirect(w, r, "/", http.StatusFound) // TODO redirect to somewhere proper
}
