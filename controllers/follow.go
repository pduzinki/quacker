package controllers

import (
	"net/http"

	"quacker/models"
)

// TODO add proper description
type FollowController struct {
	fs models.FollowService
}

// TODO add proper description
func NewFollowController(fs models.FollowService) *FollowController {
	return &FollowController{
		fs: fs,
	}
}

// TODO add proper description
func (fc *FollowController) FollowUser(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusFound)
}

// TODO add proper description
func (fc *FollowController) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusFound)
}
