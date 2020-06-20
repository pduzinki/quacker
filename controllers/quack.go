package controllers

import (
// "net/http"

// "quacker/views"
)

type QuackC struct {
	// HomeView *views.View
}

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
