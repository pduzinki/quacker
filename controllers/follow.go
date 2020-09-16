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
	fs models.FollowService
	us models.UserService
}

// NewFollowController creates new follow controller
func NewFollowController(fs models.FollowService, us models.UserService) *FollowController {
	return &FollowController{
		fs: fs,
		us: us,
	}
}

// FollowUser handles POST /{username}/follow
func (fc *FollowController) FollowUser(w http.ResponseWriter, r *http.Request) {
	loggedUser := context.GetUser(r.Context())
	params := mux.Vars(r)

	_ = loggedUser

	username, _ := params["user"]

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
	params := mux.Vars(r)

	username, _ := params["user"]

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
