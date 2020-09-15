package views

import (
	"net/http"
	"time"

	"quacker/models"
)

// Alert represents data rendered on page when something goes wrong
type Alert struct {
	Level   string
	Message string
}

// Profile wraps data needed for rendering /{username} page
type Profile struct {
	Username string
	About    string
	Exists   bool
	Self     bool
	Followed bool
	Quacks   []models.Quack
}

// Data is a helper structure for rendering data on a page
type Data struct {
	Alert *Alert
	Yield interface{}
}

// SetAlert sets value of an alert
func (d *Data) SetAlert(err error) {
	d.Alert = &Alert{
		Level:   "danger",
		Message: err.Error(),
	}
}

func saveAlert(w http.ResponseWriter, alert Alert) {
	expires := time.Now().Add(20 * time.Second)
	lvl := http.Cookie{
		Name:     "alert_level",
		Value:    alert.Level,
		Path:     "/",
		Expires:  expires,
		HttpOnly: true,
	}
	msg := http.Cookie{
		Name:     "alert_message",
		Value:    alert.Message,
		Path:     "/",
		Expires:  expires,
		HttpOnly: true,
	}

	http.SetCookie(w, &lvl)
	http.SetCookie(w, &msg)
}

func getAlert(r *http.Request) *Alert {
	lvl, err := r.Cookie("alert_level")
	if err != nil {
		return nil
	}

	msg, err := r.Cookie("alert_message")
	if err != nil {
		return nil
	}

	return &Alert{
		Level:   lvl.Value,
		Message: msg.Value,
	}
}

func clearAlert(w http.ResponseWriter) {
	lvl := http.Cookie{
		Name:     "alert_level",
		Value:    "",
		Path:     "/",
		Expires:  time.Now(),
		HttpOnly: true,
	}
	msg := http.Cookie{
		Name:     "alert_message",
		Value:    "",
		Path:     "/",
		Expires:  time.Now(),
		HttpOnly: true,
	}

	http.SetCookie(w, &lvl)
	http.SetCookie(w, &msg)
}

func RedirectWithAlert(w http.ResponseWriter, r *http.Request, url string, httpStatus int, alert Alert) {
	saveAlert(w, alert)
	http.Redirect(w, r, url, httpStatus)
}
