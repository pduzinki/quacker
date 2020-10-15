package controllers

import (
	"github.com/gorilla/mux"
	"net/http"

	"quacker/context"
	"quacker/models"
	"quacker/views"
)

// HashtagC is a controller struct responsible for handling hashtag resource
type HashtagC struct {
	HashtagView *views.View
	qs          models.QuackService
	hs          models.HashtagService
}

// NewHashtagController creates new hashtag controller
func NewHashtagController(qs models.QuackService, hs models.HashtagService) *HashtagC {
	hc := HashtagC{
		HashtagView: views.NewView("views/hashtag/hashtag.gohtml", "views/quack/quack.gohtml"),
		qs:          qs,
		hs:          hs,
	}

	return &hc
}

// ShowQuacksByHashtag ...
func (hc *HashtagC) ShowQuacksByHashtag(w http.ResponseWriter, r *http.Request) {
	var d views.Data
	vars := mux.Vars(r)

	loggedUser := context.GetUser(r.Context())
	if loggedUser == nil {
		loggedUser = &models.User{}
	}
	d.User = loggedUser

	hashtag := vars["hashtag"]

	quacks, err := hc.qs.FindByHashtag(hashtag)
	if err != nil {
		d.SetAlert(err)
		hc.HashtagView.Render(w, r, d)
		return
	}

	vQuacks := make([]views.Quack, len(quacks), len(quacks))
	for i, q := range quacks {
		vQuacks[i].Quack = q
		vQuacks[i].BelongsToLoggedUser = (loggedUser.Username == q.Username)
	}

	d.Yield = views.Profile{
		Quacks: vQuacks,
	}

	hc.HashtagView.Render(w, r, d)
}
