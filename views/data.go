package views

import (
	"quacker/models"
)

type Alert struct {
	Level   string
	Message string
}

type Profile struct {
	Username string
	About    string
	Exists   bool
	Self     bool
	Followed bool
	Quacks   []models.Quack
}

// Data
type Data struct {
	Alert *Alert
	Yield interface{}
	// User  *models.User
}

// SetAlert sets value of an alert
func (d *Data) SetAlert(err error) {
	d.Alert = &Alert{
		Level:   "danger",
		Message: err.Error(),
	}
}

// func (d *Data) SetUser(u *models.User) {
// 	d.User = u
// }

// TODO add alerts saved in cookies, for redirects
