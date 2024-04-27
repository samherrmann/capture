package main

import (
	"embed"
	"fmt"
	"io"
	"net/http"
)

//go:embed index.html
var embedFS embed.FS

func viewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := writeResponse(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func writeResponse(w http.ResponseWriter) error {
	file, err := embedFS.Open("index.html")
	if err != nil {
		return fmt.Errorf("opening HTML: %w", err)
	}
	defer file.Close()
	_, err = io.Copy(w, file)
	if err != nil {
		return fmt.Errorf("copying HTML to response: %w", err)
	}
	return nil
}
