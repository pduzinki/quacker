package controllers

import (
// "net/http"

// "quacker/views"
)

// QuackC is a controller struct responsible for handling quack resources
type QuackC struct {
	// HomeView *views.View
}

// NewQuackC creates new quack controller
func NewQuackC() *QuackC {
	qc := QuackC{
		// HomeView: views.NewView("views/quack/home.gohtml"),
	}

	return &qc
}

// GetHomepage handles GET /
// func (qc *QuackC) GetHomepage(w http.ResponseWriter, r *http.Request) {
// 	qc.HomeView.Render(w, r)
// }
