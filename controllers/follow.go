package controllers

import (
	"log"
	"net/http"

	"quacker/context"
	"quacker/models"

	"github.com/gorilla/mux"
)

// TODO add proper description
type FollowController struct {
	fs models.FollowService
	us models.UserService
}

// TODO add proper description
func NewFollowController(fs models.FollowService, us models.UserService) *FollowController {
	return &FollowController{
		fs: fs,
		us: us,
	}
}

// TODO add proper description
func (fc *FollowController) FollowUser(w http.ResponseWriter, r *http.Request) {
	loggedUser := context.GetUser(r.Context())
	params := mux.Vars(r)

	_ = loggedUser

	username, prs := params["user"]
	if prs == false {
		// TODO add some logging
	}

	userToFollow, err := fc.us.FindByUsername(username)
	if err != nil {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	_ = userToFollow

	follow := models.Follow{
		UserID:        loggedUser.ID,
		FollowsUserID: userToFollow.ID,
	}

	err = fc.fs.Create(&follow)
	if err != nil {
		log.Println("Failed to create follow relation: ", err)
	}

	http.Redirect(w, r, "/"+userToFollow.Username, http.StatusFound)
}

// TODO add proper description
func (fc *FollowController) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusFound)
}
