package controllers

import (
	"net/http"

	"quacker/views"
)

type FeedC struct {
	HomeView *views.View
}

func NewFeedC() *FeedC {
	fc := FeedC{
		HomeView: views.NewView("views/feed/home.gohtml"),
	}

	return &fc
}

// GetHomepage handles GET /
func (fc *FeedC) GetHomepage(w http.ResponseWriter, r *http.Request) {
	fc.HomeView.Render(w, r)
}
