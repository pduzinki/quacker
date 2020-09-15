package views

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
)

// View wraps templates and everything necessary to render a page
type View struct {
	Template *template.Template
	// Layout   string
}

// NewView creates View instance
func NewView(files ...string) *View {
	files = append(files, "views/layouts/alert.gohtml",
		"views/layouts/footer.gohtml",
		"views/layouts/header.gohtml",
		"views/layouts/main.gohtml",
		"views/layouts/navbar.gohtml")

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	v := View{
		Template: t,
	}

	return &v
}

// Render prepares an http response to render a page
func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	var vd Data

	switch d := data.(type) {
	case Data:
		vd = d
	default:
		vd = Data{
			Yield: data,
		}
	}

	alert := getAlert(r)
	if alert != nil {
		vd.Alert = alert
		clearAlert(w)
	}

	var buf bytes.Buffer
	err := v.Template.ExecuteTemplate(&buf, "main", vd)
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}
