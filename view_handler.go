package main

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/samherrmann/capture/cookies"
)

//go:embed view.html
var viewFS embed.FS

type ViewData struct {
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

func newViewHandler() (http.Handler, error) {
	tpl, err := template.ParseFS(viewFS, "view.html")
	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		status, err := cookies.GetStatus(r)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				writeView(w, http.StatusOK, tpl, nil)
				return
			}
			status = &cookies.Status{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("status cookie: %s", err),
			}
		}

		statusLevel := StatusLevelError
		if status.Code < 400 {
			statusLevel = StatusLevelSuccess
		}

		// Ensure cookies are cleared regardless if there was an error or not.
		cookies.ClearStatus(w)

		viewData := &ViewData{
			Status: &Status{
				Level:   statusLevel,
				Message: status.Message,
			},
		}

		writeView(w, status.Code, tpl, viewData)
	}), nil
}

// writeView writes the view to the response w.
func writeView(w http.ResponseWriter, code int, tpl *template.Template, data *ViewData) {
	w.WriteHeader(code)
	if err := tpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
