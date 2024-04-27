package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fileUploadHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code, err := copyFormFile(r)
		if err != nil {
			http.Error(w, err.Error(), code)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}

func copyFormFile(r *http.Request) (int, error) {
	if err := r.ParseMultipartForm(200 << 20); err != nil {
		err := fmt.Errorf("parsing multi-part form: %w", err)
		return http.StatusBadRequest, err
	}

	srcFile, handler, err := r.FormFile("file")
	if err != nil {
		err := fmt.Errorf("reading form file: %w", err)
		return http.StatusBadRequest, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(handler.Filename)
	if err != nil {
		err := fmt.Errorf("creating destination file: %w", err)
		return http.StatusInternalServerError, err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		err := fmt.Errorf("copying source file to destination file: %w", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
