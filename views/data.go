package views

import "net/http"

type Alert struct {
	Level   string
	Message string
}

// Data
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

// TODO add alerts saved in cookies, for redirects
