package controllers

import (
// "net/http"

// "quacker/views"
)

// HashtagC is a controller struct responsible for handling hashtag resource
type HashtagC struct {
	// HomeView *views.View
}

// NewHashtagC creates new hashtag controller
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
