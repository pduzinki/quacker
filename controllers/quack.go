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
	fs          models.FollowService
}

// NewQuackController creates new quack controller
func NewQuackController(qs models.QuackService, us models.UserService, fs models.FollowService) *QuackController {
	qc := QuackController{
		HomeView:    views.NewView("views/quack/home.gohtml"),
		ProfileView: views.NewView("views/quack/profile.gohtml", "views/follow/follow.gohtml", "views/follow/unfollow.gohtml"),
		qs:          qs,
		us:          us,
		fs:          fs,
	}

	return &qc
}

// GetHome handles GET /home
func (qc *QuackController) GetHome(w http.ResponseWriter, r *http.Request) {
	user := context.GetUser(r.Context())

	var d views.Data
	d.User = user

	qc.HomeView.Render(w, r, d)
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

	user := context.GetUser(r.Context())
	vd.User = user

	username, _ := params["user"]

	// check if user with such an username exists, get user
	user, err := qc.us.FindByUsername(username)
	if err == models.ErrRecordNotFound {
		vd.Yield = views.Profile{
			Exists:   false,
			Username: username,
		}
		qc.ProfileView.Render(w, r, vd)
		return
	} else if err != nil {
		vd.SetAlert(err)
		qc.ProfileView.Render(w, r, vd)
		return
	}

	loggedUser := &models.User{}
	cookie, _ := r.Cookie("remember_token")
	if cookie != nil {
		loggedUser, _ = qc.us.FindByRememberToken(cookie.Value)
	}

	var self bool
	var followed bool
	if loggedUser.ID == user.ID {
		self = true
	} else {
		if loggedUser != nil {
			// TODO replace with call to FindByIDs

			follows, _ := qc.fs.FindByUserID(loggedUser.ID)

			for _, follow := range follows {
				if follow.FollowsUserID == user.ID {
					followed = true
				}
			}
		}
	}

	quacks, err := qc.qs.FindByUserID(user.ID)
	if err != nil {
		vd.SetAlert(err)
		qc.ProfileView.Render(w, r, vd)
		return
	}

	vd.Yield = views.Profile{
		Username: user.Username,
		About:    user.About,
		Exists:   true,
		Self:     self,
		Followed: followed,
		Quacks:   quacks,
	}

	// render page
	qc.ProfileView.Render(w, r, vd)
}
