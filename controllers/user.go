package controllers

import (
	// "fmt"
	"net/http"

	"github.com/gorilla/mux"

	"quacker/models"
	"quacker/token"
	"quacker/views"
)

// UserController is responsible for handling user resources
type UserController struct {
	HomepageView *views.View
	LoginView    *views.View
	SignupView   *views.View
	UsernameView *views.View
	us           models.UserService
}

// NewUserController creates user controller instance
func NewUserController(us models.UserService) *UserController {
	uc := UserController{
		HomepageView: views.NewView("views/user/homepage.gohtml"),
		LoginView:    views.NewView("views/user/login.gohtml"),
		SignupView:   views.NewView("views/user/signup.gohtml"),
		UsernameView: views.NewView("views/user/user.gohtml"),
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

	err = uc.signIn(w, &user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// TODO redirect to /<username>
	http.Redirect(w, r, "/", http.StatusFound)
}

// signIn creates a cookie with remember token for the given user
func (uc *UserController) signIn(w http.ResponseWriter, u *models.User) error {
	if u.RememberToken == "" {
		token, err := token.GenerateRememberToken()
		if err != nil {
			return err
		}
		u.RememberToken = token
		err = uc.us.Update(u)
		if err != nil {
			return err
		}
	}

	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    u.RememberToken,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	return nil
}

// GetUser handles GET /{username}
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	// read username from the url
	params := mux.Vars(r)

	username, prs := params["user"]
	if prs == false {
		// TODO add some logging
	}

	// check if user with such an username exists, get user
	user, err := uc.us.FindByUsername(username)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	_ = user

	// fill data for template
	// render page

	var d views.Data
	d.User = "obi-wan kenobi"

	uc.UsernameView.Render(w, r, d)
}
