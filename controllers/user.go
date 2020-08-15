package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"quacker/models"
	"quacker/token"
	"quacker/views"
)

// UserController is responsible for handling user resources
type UserController struct {
	WelcomeView  *views.View
	LoginView    *views.View
	SignupView   *views.View
	UsernameView *views.View
	HomeView     *views.View
	us           models.UserService
}

// NewUserController creates user controller instance
func NewUserController(us models.UserService) *UserController {
	uc := UserController{
		WelcomeView:  views.NewView("views/user/welcome.gohtml"),
		LoginView:    views.NewView("views/user/login.gohtml"),
		SignupView:   views.NewView("views/user/signup.gohtml"),
		UsernameView: views.NewView("views/user/user.gohtml"),
		HomeView:     views.NewView("views/user/home.gohtml"),
		us:           us,
	}

	return &uc
}

// GetWelcome handles GET /
func (uc *UserController) GetWelcome(w http.ResponseWriter, r *http.Request) {
	uc.WelcomeView.Render(w, r, nil)
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
		return
	}

	err = uc.signIn(w, &user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// TODO redirect to /<username>
	http.Redirect(w, r, "/", http.StatusFound)
}

// GetHome handles GET /home
func (uc *UserController) GetHome(w http.ResponseWriter, r *http.Request) {
	// TODO add proper /home page
	uc.HomeView.Render(w, r, nil)
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
	var vd views.Data
	params := mux.Vars(r)

	username, prs := params["user"]
	if prs == false {
		// TODO add some logging
	}

	// check if user with such an username exists, get user
	user, err := uc.us.FindByUsername(username)
	if err == models.ErrRecordNotFound {
		// create default user to be returned
		user = &models.User{
			Username: username,
			About:    "user doesn't exist.",
		}
	} else if err != nil {
		vd.SetAlert(err)
	}

	// fill data for template
	vd.SetUser(user)

	// render page
	uc.UsernameView.Render(w, r, vd)
}

// GetNewQuack handles GET /quack
func (uc *UserController) GetNewQuack(w http.ResponseWriter, r *http.Request) {
	uc.HomeView.Render(w, r, nil)
}

// PostNewQuack handles POST /quack
func (uc *UserController) PostNewQuack(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// CookieTest handles GET /cookietest, this function is for testing only
func (uc *UserController) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("remember_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := uc.us.FindByRememberToken(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, user)
}
