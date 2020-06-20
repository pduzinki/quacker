package controllers

import (
// "net/http"

// "quacker/views"
)

type HashtagC struct {
	// HomeView *views.View
}

func NewHashtagC() *HashtagC {
	hc := HashtagC{
		// HomeView: views.NewView("views/hashtag/home.gohtml"),
	}

	return &hc
}

// GetHomepage handles GET /
// func (hc *HashtagC) GetHomepage(w http.ResponseWriter, r *http.Request) {
// 	hc.HomeView.Render(w, r)
// }
