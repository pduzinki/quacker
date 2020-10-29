package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"quacker/context"
	"quacker/match"
	"quacker/models"
	"quacker/truncate"
	"quacker/views"
)

// QuackController is a controller struct responsible for handling quack resources
type QuackController struct {
	HomeView    *views.View
	ProfileView *views.View
	QuackView   *views.View
	TagView     *views.View
	qs          models.QuackService
	us          models.UserService
	fs          models.FollowService
	hs          models.TagService
	tagRegex    *regexp.Regexp
	atRegex     *regexp.Regexp
}

// NewQuackController creates new quack controller
func NewQuackController(qs models.QuackService, us models.UserService,
	fs models.FollowService, hs models.TagService) *QuackController {
	qc := QuackController{
		HomeView: views.NewView("views/quack/home.gohtml", "views/quack/quack.gohtml"),
		ProfileView: views.NewView("views/quack/profile.gohtml",
			"views/quack/quack.gohtml",
			"views/follow/follow.gohtml",
			"views/follow/unfollow.gohtml"),
		QuackView: views.NewView("views/quack/quackpage.gohtml"),
		TagView:   views.NewView("views/tag/tag.gohtml", "views/quack/quack.gohtml"),
		qs:        qs,
		us:        us,
		fs:        fs,
		hs:        hs,
		tagRegex:  regexp.MustCompile(match.Tag()),
		atRegex:   regexp.MustCompile(match.At()),
	}

	return &qc
}

// GetHome handles GET /home
func (qc *QuackController) GetHome(w http.ResponseWriter, r *http.Request) {
	loggedUser := context.GetUser(r.Context())

	var d views.Data
	d.User = loggedUser
	d.Yield = views.Profile{}

	// get all followed users
	follows, err := qc.fs.FindByUserID(loggedUser.ID)
	if err != nil {
		d.SetAlert(err)
		qc.HomeView.Render(w, r, d)
		return
	}

	// extract followed users IDs
	followsIDs := make([]uint, len(follows), len(follows))
	for i, f := range follows {
		followsIDs[i] = f.FollowsUserID
	}
	log.Println("fine")

	// add yourself to IDs, to see your quacks on the quack board
	followsIDs = append(followsIDs, loggedUser.ID)

	// query db for all quacks
	quacks, err := qc.qs.FindByMultipleUserIDs(followsIDs)
	if err != nil {
		d.SetAlert(err)
		qc.HomeView.Render(w, r, d)
		return
	}

	vQuacks := make([]views.Quack, len(quacks), len(quacks))
	for i, q := range quacks {
		vQuacks[i].Quack = q
		vQuacks[i].QuackTextParts = qc.ParseQuackText(q.Text)
		vQuacks[i].BelongsToLoggedUser = (loggedUser.Username == q.Username)
	}

	d.Yield = views.Profile{
		Quacks: vQuacks,
	}

	qc.HomeView.Render(w, r, d)
}

// NewQuack handles POST /home (i.e. posting new quacks)
func (qc *QuackController) NewQuack(w http.ResponseWriter, r *http.Request) {
	var form quackForm

	err := parseForm(r, &form)
	if err != nil {
		alert := views.Alert{
			Level:   "danger",
			Message: err.Error(),
		}
		views.RedirectWithAlert(w, r, "/home", http.StatusFound, alert)
		return
	}

	user := context.GetUser(r.Context())

	quack := models.Quack{
		UserID: user.ID,
		Text:   form.Text,
	}

	err = qc.qs.Create(&quack)
	if err != nil {
		alert := views.Alert{
			Level:   "danger",
			Message: err.Error(),
		}
		views.RedirectWithAlert(w, r, "/home", http.StatusFound, alert)
		return
	}

	tagStrings := qc.hs.ParseTags(quack.Text)
	for _, tagString := range tagStrings {
		tag := models.Tag{
			Text:    tagString,
			QuackID: quack.ID,
		}

		err = qc.hs.Create(&tag)
		if err != nil {
			log.Println("Failed to create a tag.")
			// unfortunate, but let's keep going
		}
	}

	http.Redirect(w, r, "/home", http.StatusFound)
}

// GetProfile handles GET /{username}
func (qc *QuackController) GetProfile(w http.ResponseWriter, r *http.Request) {
	// read username from the url
	var vd views.Data
	vars := mux.Vars(r)

	loggedUser := context.GetUser(r.Context())
	vd.User = loggedUser

	username, _ := vars["user"]

	// check if user with such an username exists, get user
	user, err := qc.us.FindByUsername(username)
	if err == models.ErrRecordNotFound {
		vd.Yield = views.Profile{
			Exists:   false,
			Username: username,
		}
		qc.ProfileView.Render(w, r, vd)
		return
	} else if err != nil {
		vd.SetAlert(err)
		qc.ProfileView.Render(w, r, vd)
		return
	}

	var self bool
	var followed bool
	if loggedUser != nil {
		if loggedUser.ID == user.ID {
			self = true
		} else {
			_, err := qc.fs.FindByIDs(loggedUser.ID, user.ID)
			if err == nil {
				followed = true
			}
		}
	}

	quacks, err := qc.qs.FindByUserID(user.ID)
	if err != nil {
		vd.SetAlert(err)
		qc.ProfileView.Render(w, r, vd)
		return
	}

	vQuacks := make([]views.Quack, len(quacks), len(quacks))
	for i, q := range quacks {
		vQuacks[i].Quack = q
		vQuacks[i].QuackTextParts = qc.ParseQuackText(q.Text)
		vQuacks[i].BelongsToLoggedUser = (loggedUser != nil) && (loggedUser.Username == q.Username)
	}

	vd.Yield = views.Profile{
		Username: user.Username,
		About:    user.About,
		Exists:   true,
		Self:     self,
		Followed: followed,
		Quacks:   vQuacks,
	}

	// render page
	qc.ProfileView.Render(w, r, vd)
}

// GetQuack handles GET /{user}/quacks/{id}
func (qc *QuackController) GetQuack(w http.ResponseWriter, r *http.Request) {
	loggedUser := context.GetUser(r.Context())
	vars := mux.Vars(r)

	var d views.Data
	d.User = loggedUser

	username, _ := vars["user"]
	id64, _ := strconv.ParseUint(vars["id"], 10, 32)
	id := uint(id64)

	user, err := qc.us.FindByUsername(username)
	if err != nil {
		user = &models.User{}
	}

	quack, err := qc.qs.FindByID(id)
	if err != nil {
		d.SetAlert(err)
		qc.QuackView.Render(w, r, d)
		return
	}

	if quack.UserID != user.ID {
		url := "/" + quack.Username + "/quack/" + strconv.Itoa(int(quack.ID))
		http.Redirect(w, r, url, http.StatusFound)
		return
	}

	vQuack := views.Quack{
		Quack:               *quack,
		QuackTextParts:      qc.ParseQuackText(quack.Text),
		BelongsToLoggedUser: (loggedUser.ID == quack.UserID),
	}

	d.Yield = vQuack
	qc.QuackView.Render(w, r, d)
}

// DeleteQuack handles POST /{user}/quacks/{id}/delete
func (qc *QuackController) DeleteQuack(w http.ResponseWriter, r *http.Request) {
	loggedUser := context.GetUser(r.Context())
	vars := mux.Vars(r)

	// username := vars["user"] // not used
	id64, _ := strconv.ParseUint(vars["id"], 10, 32)
	id := uint(id64)

	quack, err := qc.qs.FindByID(id)
	if err != nil {
		alert := views.Alert{
			Level:   "danger",
			Message: "Quack not found.",
		}
		views.RedirectWithAlert(w, r, "/home", http.StatusFound, alert)
	}

	if quack.UserID != loggedUser.ID {
		alert := views.Alert{
			Level:   "danger",
			Message: "You're not authorized to delete this quack.",
		}
		views.RedirectWithAlert(w, r, "/home", http.StatusFound, alert)
	}

	err = qc.qs.Delete(id)
	if err != nil {
		alert := views.Alert{
			Level:   "danger",
			Message: "Failed to delete quack.",
		}
		views.RedirectWithAlert(w, r, "/home", http.StatusFound, alert)
	}
	// TODO delete entries from tags table

	http.Redirect(w, r, "/home", http.StatusFound)
}

// ShowQuacksByTag handles GET /tags/{tag}
func (qc *QuackController) ShowQuacksByTag(w http.ResponseWriter, r *http.Request) {
	var d views.Data
	vars := mux.Vars(r)

	loggedUser := context.GetUser(r.Context())
	d.User = loggedUser

	tag := vars["tag"]
	// TODO verify that tag is proper

	quacks, err := qc.qs.FindByTag(tag)
	if err != nil {
		d.SetAlert(err)
		qc.TagView.Render(w, r, d)
		return
	}

	vQuacks := make([]views.Quack, len(quacks), len(quacks))
	for i, q := range quacks {
		vQuacks[i].Quack = q
		vQuacks[i].QuackTextParts = qc.ParseQuackText(q.Text)
		vQuacks[i].BelongsToLoggedUser = (loggedUser != nil) && (loggedUser.Username == q.Username)
	}

	d.Yield = views.Profile{
		Quacks: vQuacks,
	}

	qc.TagView.Render(w, r, d)
}

// ParseQuackText parses quackText and wraps #tags and @ats into template.HTML
func (qc *QuackController) ParseQuackText(quackText string) []interface{} {
	// TODO this feels like it should be a part of views package instead
	quackTextParts := make([]interface{}, 0)
	words := strings.Split(quackText, " ")

	for _, word := range words {
		if qc.tagRegex.MatchString(word) {
			word = truncate.WithoutFirstRune(word)
			link := fmt.Sprintf(`<a href="/tags/%v">#%v</a>`, word, word)
			quackTextParts = append(quackTextParts, template.HTML(link))
		} else if qc.atRegex.MatchString(word) {
			word = truncate.WithoutFirstRune(word)
			link := fmt.Sprintf(`<a href="/%v">@%v</a>`, word, word)
			quackTextParts = append(quackTextParts, template.HTML(link))
		} else {
			quackTextParts = append(quackTextParts, word)
		}
	}

	return quackTextParts
}
