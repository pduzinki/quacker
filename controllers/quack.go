package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"quacker/context"
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
		ProfileView: views.NewView("views/quack/profile.gohtml", "views/follow/follow.gohtml", "views/follow/unfollow.gohtml"),
		qs:          qs,
		us:          us,
	}

	return &qc
}

// GetHome handles GET /home
func (qc *QuackController) GetHome(w http.ResponseWriter, r *http.Request) {
	qc.HomeView.Render(w, r, nil)
}

// NewQuack handles POST /home (i.e. posting new quacks)
func (qc *QuackController) NewQuack(w http.ResponseWriter, r *http.Request) {
	var form quackForm
	var d views.Data

	err := parseForm(r, &form)
	if err != nil {
		d.SetAlert(err)
		qc.HomeView.Render(w, r, d)
		return
	}

	user := context.GetUser(r.Context())

	quack := models.Quack{
		UserID: user.ID,
		Text:   form.Text,
	}

	err = qc.qs.Create(&quack)
	if err != nil {
		d.SetAlert(err)
		qc.HomeView.Render(w, r, d)
		return
	}

	qc.HomeView.Render(w, r, d)
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
		qc.ProfileView.Render(w, r, vd)
		return
	}

	// fill data for template
	vd.SetUser(user)

	quacks, err := qc.qs.FindByUserID(user.ID)
	if err != nil {
		vd.SetAlert(err)
		qc.ProfileView.Render(w, r, vd)
		return
	}

	// user didn't quack anything
	if len(quacks) == 0 {
		// TODO this isn't really elegant, consider refactoring profile.gohtml template
		nilQuack := models.Quack{
			Text: "user didn't quack anything yet.",
		}
		vd.Yield = []models.Quack{nilQuack}
		qc.ProfileView.Render(w, r, vd)
		return
	}

	vd.Yield = quacks

	// render page
	qc.ProfileView.Render(w, r, vd)
}
