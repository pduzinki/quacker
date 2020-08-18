package controllers

import (
	"net/http"

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

// GetProfile handles GET /{:username}
func (qc *QuackController) GetProfile(w http.ResponseWriter, r *http.Request) {
	qc.ProfileView.Render(w, r, nil)
}

// NewQuack handles POST /home
func (qc *QuackController) NewQuack(w http.ResponseWriter, r *http.Request) {
	// TODO
}
