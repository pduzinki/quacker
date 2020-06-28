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
func NewView(file string) *View {
	t, err := template.ParseFiles("views/layouts/page.gohtml", "views/layouts/navbar.gohtml", file)
	if err != nil {
		panic(err)
	}

	v := View{
		Template: t,
	}

	return &v
}

// Render prepares an http response to render a page
func (v *View) Render(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	var buf bytes.Buffer
	err := v.Template.ExecuteTemplate(&buf, "page", nil)
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}
