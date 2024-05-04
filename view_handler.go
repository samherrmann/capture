package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/samherrmann/capture/cookies"
	"github.com/samherrmann/capture/view"
)

func newViewHandler() (http.Handler, error) {
	tpl, err := view.NewTemplate()
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

		// Get status level based on status code.
		statusLevel := view.StatusLevelError
		if status.Code < 400 {
			statusLevel = view.StatusLevelSuccess
		}

		viewData := &view.Data{
			Status: &view.Status{
				Level:   statusLevel,
				Message: status.Message,
			},
		}

		writeView(w, status.Code, tpl, viewData)
	}), nil
}

// writeView writes the view to the response w.
func writeView(w http.ResponseWriter, code int, tpl *template.Template, data *view.Data) {
	cookies.ClearStatus(w)
	w.WriteHeader(code)
	if err := tpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
