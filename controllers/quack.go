package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"quacker/context"
	"quacker/models"
	"quacker/views"
)

// QuackController is a controller struct responsible for handling quack resources
type QuackController struct {
	HomeView    *views.View
	ProfileView *views.View
	QuackView   *views.View
	qs          models.QuackService
	us          models.UserService
	fs          models.FollowService
}

// NewQuackController creates new quack controller
func NewQuackController(qs models.QuackService, us models.UserService, fs models.FollowService) *QuackController {
	qc := QuackController{
		HomeView: views.NewView("views/quack/home.gohtml", "views/quack/quack.gohtml"),
		ProfileView: views.NewView("views/quack/profile.gohtml",
			"views/quack/quack.gohtml",
			"views/follow/follow.gohtml",
			"views/follow/unfollow.gohtml"),
		QuackView: views.NewView("views/quack/quackpage.gohtml"),
		qs:        qs,
		us:        us,
		fs:        fs,
	}

	return &qc
}

// GetHome handles GET /home
func (qc *QuackController) GetHome(w http.ResponseWriter, r *http.Request) {
	user := context.GetUser(r.Context())

	var d views.Data
	d.User = user
	d.Yield = views.Profile{}

	// get all followed users
	follows, err := qc.fs.FindByUserID(user.ID)
	if err != nil {
		d.SetAlert(err)
		qc.HomeView.Render(w, r, d)
		return
	}

	// extract followed users IDs
	followsIDs := make([]uint, len(follows), len(follows))
	for i, f := range follows {
		followsIDs[i] = f.FollowsUserID
	}
	log.Println("fine")

	// add yourself to IDs, to see your quacks on the quack board
	followsIDs = append(followsIDs, user.ID)

	// query db for all quacks
	quacks, err := qc.qs.FindByMultipleUserIDs(followsIDs)
	if err != nil {
		d.SetAlert(err)
		qc.HomeView.Render(w, r, d)
		return
	}

	d.Yield = views.Profile{
		Quacks: quacks,
	}

	qc.HomeView.Render(w, r, d)
}

// NewQuack handles POST /home (i.e. posting new quacks)
func (qc *QuackController) NewQuack(w http.ResponseWriter, r *http.Request) {
	var form quackForm

	err := parseForm(r, &form)
	if err != nil {
		alert := views.Alert{
			Level:   "danger",
			Message: err.Error(),
		}
		views.RedirectWithAlert(w, r, "/home", http.StatusFound, alert)
		return
	}

	user := context.GetUser(r.Context())

	quack := models.Quack{
		UserID: user.ID,
		Text:   form.Text,
	}

	err = qc.qs.Create(&quack)
	if err != nil {
		alert := views.Alert{
			Level:   "danger",
			Message: err.Error(),
		}
		views.RedirectWithAlert(w, r, "/home", http.StatusFound, alert)
		return
	}

	http.Redirect(w, r, "/home", http.StatusFound)
}

// GetProfile handles GET /{username}
func (qc *QuackController) GetProfile(w http.ResponseWriter, r *http.Request) {
	// read username from the url
	var vd views.Data
	params := mux.Vars(r)

	loggedUser := context.GetUser(r.Context())
	vd.User = loggedUser

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

	var self bool
	var followed bool
	if loggedUser != nil {
		if loggedUser.ID == user.ID {
			self = true
		} else {
			_, err := qc.fs.FindByIDs(loggedUser.ID, user.ID)
			if err == nil {
				followed = true
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

// GetQuack handles GET /{user}/quacks/{id}
func (qc *QuackController) GetQuack(w http.ResponseWriter, r *http.Request) {
	loggedUser := context.GetUser(r.Context())
	vars := mux.Vars(r)

	var d views.Data
	d.User = loggedUser

	username, _ := vars["user"]
	id64, _ := strconv.ParseUint(vars["id"], 10, 32)
	id := uint(id64)

	user, err := qc.us.FindByUsername(username)
	if err != nil {
		user = &models.User{}
	}

	quack, err := qc.qs.FindByID(id)
	if err != nil {
		d.SetAlert(err)
		qc.QuackView.Render(w, r, d)
		return
	}

	if quack.UserID != user.ID {
		url := "/" + quack.Username + "/quack/" + strconv.Itoa(int(quack.ID))
		http.Redirect(w, r, url, http.StatusFound)
		return
	}

	d.Yield = quack
	qc.QuackView.Render(w, r, d)
}
