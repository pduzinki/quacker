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
	uc.HomepageView.Render(w, r)
}

// GetLogin handles GET /login
func (uc *UserController) GetLogin(w http.ResponseWriter, r *http.Request) {
	uc.LoginView.Render(w, r)
}

// PostLogin handles POST /login
func (uc *UserController) PostLogin(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// GetSignup handles GET /signup
func (uc *UserController) GetSignup(w http.ResponseWriter, r *http.Request) {
	uc.SignupView.Render(w, r)
}

// PostSignup handles POST /signup
func (uc *UserController) PostSignup(w http.ResponseWriter, r *http.Request) {
	var form signupForm

	err := parseForm(r, &form)
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username: form.Username,
		Email:    form.Email,
		Password: form.Password,
	}

	uc.us.Create(&user)

	// fmt.Print(form)
}
