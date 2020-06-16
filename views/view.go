package views

import (
	"html/template"
)

type View struct {
	Template *template.Template
	Layout   string
}
