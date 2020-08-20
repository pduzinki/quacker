package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"quacker/models"
	"quacker/views"
)

// QuackController is a controller struct responsible for handling quack resources
type QuackController struct {
	HomeView    *views.View
	ProfileView *views.View
	qs          models.QuackService
	us          models.UserService
}

// NewQuackController creates new quack controller
func NewQuackController(qs models.QuackService, us models.UserService) *QuackController {
	qc := QuackController{
		HomeView:    views.NewView("views/quack/home.gohtml"),
		ProfileView: views.NewView("views/quack/profile.gohtml"),
		qs:          qs,
		us:          us,
	}

	return &qc
}

// GetHome handles GET /home
func (qc *QuackController) GetHome(w http.ResponseWriter, r *http.Request) {
	qc.HomeView.Render(w, r, nil)
}

// GetProfile handles GET /{username}
func (qc *QuackController) GetProfile(w http.ResponseWriter, r *http.Request) {
	// read username from the url
	var vd views.Data
	params := mux.Vars(r)

	username, prs := params["user"]
	if prs == false {
		// TODO add some logging
	}

	// check if user with such an username exists, get user
	user, err := qc.us.FindByUsername(username)
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
	qc.ProfileView.Render(w, r, vd)

	// qc.ProfileView.Render(w, r, nil)
}

// NewQuack handles POST /home
func (qc *QuackController) NewQuack(w http.ResponseWriter, r *http.Request) {
	// TODO
}
