package view

import (
	"embed"
	"html/template"
)

//go:embed view.html
var viewFS embed.FS

type Data struct {
	Status *Status
}

type Status struct {
	Level   StatusLevel
	Message string
}

type StatusLevel string

const (
	StatusLevelSuccess StatusLevel = "success"
	StatusLevelError   StatusLevel = "error"
)

func NewTemplate() (*template.Template, error) {
	return template.ParseFS(viewFS, "view.html")
}
