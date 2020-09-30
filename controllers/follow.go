package controllers

import (
	"net/http"

	"quacker/context"
	"quacker/models"
	"quacker/views"

	"github.com/gorilla/mux"
)

// FollowController is a controller struct responsible for handling follow resources
type FollowController struct {
	fs            models.FollowService
	us            models.UserService
	followingView *views.View
	followersView *views.View
}

// NewFollowController creates new follow controller
func NewFollowController(fs models.FollowService, us models.UserService) *FollowController {
	return &FollowController{
		fs:            fs,
		us:            us,
		followingView: views.NewView("views/follow/following.gohtml", "views/follow/followuser.gohtml"),
		followersView: views.NewView("views/follow/followers.gohtml", "views/follow/followuser.gohtml"),
	}
}

// FollowUser handles POST /{username}/follow
func (fc *FollowController) FollowUser(w http.ResponseWriter, r *http.Request) {
	loggedUser := context.GetUser(r.Context())
	vars := mux.Vars(r)

	_ = loggedUser

	username, _ := vars["user"]

	userToFollow, err := fc.us.FindByUsername(username)
	if err != nil {
		alert := views.Alert{
			Level:   "danger",
			Message: err.Error(),
		}
		views.RedirectWithAlert(w, r, "/"+userToFollow.Username, http.StatusFound, alert)
		return
	}

	follow := models.Follow{
		UserID:        loggedUser.ID,
		FollowsUserID: userToFollow.ID,
	}

	err = fc.fs.Create(&follow)
	if err != nil {
		alert := views.Alert{
			Level:   "danger",
			Message: err.Error(),
		}
		views.RedirectWithAlert(w, r, "/"+userToFollow.Username, http.StatusFound, alert)
		return
	}

	http.Redirect(w, r, "/"+userToFollow.Username, http.StatusFound)
}

// UnfollowUser handles POST /{username}/unfollow
func (fc *FollowController) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	loggedUser := context.GetUser(r.Context())
	vars := mux.Vars(r)

	username, _ := vars["user"]

	userToUnfollow, err := fc.us.FindByUsername(username)
	if err != nil {
		alert := views.Alert{
			Level:   "danger",
			Message: err.Error(),
		}
		views.RedirectWithAlert(w, r, "/"+userToUnfollow.Username, http.StatusFound, alert)
		return
	}

	follow, err := fc.fs.FindByIDs(loggedUser.ID, userToUnfollow.ID)
	if err != nil {
		alert := views.Alert{
			Level:   "danger",
			Message: err.Error(),
		}
		views.RedirectWithAlert(w, r, "/"+userToUnfollow.Username, http.StatusFound, alert)
		return
	}

	err = fc.fs.Delete(follow.ID)
	if err != nil {
		alert := views.Alert{
			Level:   "danger",
			Message: err.Error(),
		}
		views.RedirectWithAlert(w, r, "/"+userToUnfollow.Username, http.StatusFound, alert)
		return
	}

	http.Redirect(w, r, "/"+userToUnfollow.Username, http.StatusFound)
}

// Following handles GET /{username}/following
func (fc *FollowController) Following(w http.ResponseWriter, r *http.Request) {
	loggedUser := context.GetUser(r.Context())

	var d views.Data
	d.User = loggedUser

	vars := mux.Vars(r)
	username, _ := vars["user"]

	user, err := fc.us.FindByUsername(username)
	if err != nil {
		d.SetAlert(err)
		fc.followingView.Render(w, r, d)
		return
	}

	follows, err := fc.fs.FindByUserID(user.ID)
	if err == models.ErrRecordNotFound {
		fc.followingView.Render(w, r, d)
		return
	} else if err != nil {
		d.SetAlert(err)
		fc.followingView.Render(w, r, d)
		return
	}

	d.Yield = follows
	fc.followingView.Render(w, r, d)
}

// Followers handles GET /{username}/followers
func (fc *FollowController) Followers(w http.ResponseWriter, r *http.Request) {
	loggedUser := context.GetUser(r.Context())

	var d views.Data
	d.User = loggedUser

	vars := mux.Vars(r)
	username, _ := vars["user"]

	user, err := fc.us.FindByUsername(username)
	if err != nil {
		d.SetAlert(err)
		fc.followersView.Render(w, r, d)
		return
	}

	follows, err := fc.fs.FindByFollowsUserID(user.ID)
	if err == models.ErrRecordNotFound {
		fc.followersView.Render(w, r, d)
		return
	} else if err != nil {
		d.SetAlert(err)
		fc.followersView.Render(w, r, d)
		return
	}

	d.Yield = follows

	fc.followersView.Render(w, r, d)
}
