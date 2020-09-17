package controllers

import (
	"fmt"
	"net/http"
	"time"

	"quacker/context"
	"quacker/models"
	"quacker/token"
	"quacker/views"
)

// UserController is responsible for handling user resources
type UserController struct {
	WelcomeView *views.View
	LoginView   *views.View
	SignupView  *views.View
	us          models.UserService
}

// NewUserController creates user controller instance
func NewUserController(us models.UserService) *UserController {
	uc := UserController{
		WelcomeView: views.NewView("views/user/welcome.gohtml"),
		LoginView:   views.NewView("views/user/login.gohtml"),
		SignupView:  views.NewView("views/user/signup.gohtml"),
		us:          us,
	}

	return &uc
}

// GetWelcome handles GET /
func (uc *UserController) GetWelcome(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("remember_token")
	if err != nil {
		uc.WelcomeView.Render(w, r, nil)
		return
	}

	_, err = uc.us.FindByRememberToken(cookie.Value)
	if err != nil {
		uc.WelcomeView.Render(w, r, nil)
		return
	}

	// user is logged in, redirect to "/home"
	http.Redirect(w, r, "/home", http.StatusFound)
}

// GetLogin handles GET /login
func (uc *UserController) GetLogin(w http.ResponseWriter, r *http.Request) {
	user := context.GetUser(r.Context())
	if user != nil {
		http.Redirect(w, r, "/home", http.StatusFound)
	}

	uc.LoginView.Render(w, r, nil)
}

// PostLogin handles POST /login
func (uc *UserController) PostLogin(w http.ResponseWriter, r *http.Request) {
	var form loginForm
	var d views.Data

	err := parseForm(r, &form)
	if err != nil {
		d.SetAlert(err)
		uc.LoginView.Render(w, r, d)
		return
	}

	user, err := uc.us.Authenticate(form.Login, form.Password)
	if err != nil {
		d.SetAlert(err)
		uc.LoginView.Render(w, r, d)
		return
	}

	err = uc.signIn(w, user)
	if err != nil {
		d.SetAlert(err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/home", http.StatusFound)
}

// GetSignup handles GET /signup
func (uc *UserController) GetSignup(w http.ResponseWriter, r *http.Request) {
	user := context.GetUser(r.Context())
	if user != nil {
		http.Redirect(w, r, "/home", http.StatusFound)
	}

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
		d.SetAlert(err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/home", http.StatusFound)
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

// PostLogout handles POST /logout
func (uc *UserController) PostLogout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	user := context.GetUser(r.Context())
	rememberToken, _ := token.GenerateRememberToken()
	user.RememberToken = rememberToken
	uc.us.Update(user)
	alert := views.Alert{
		Level:   "success",
		Message: "You logged out.",
	}
	views.RedirectWithAlert(w, r, "/", http.StatusFound, alert)
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
