package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func newFileUploadHandler(dst string) (http.Handler, error) {
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		err := fmt.Errorf("making destination directory: %w", err)
		return nil, err
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code, err := copyFormFile(r, dst)
		if err != nil {
			http.Error(w, err.Error(), code)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}), nil
}

func copyFormFile(r *http.Request, dst string) (int, error) {
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

	dstFile, err := os.Create(filepath.Join(dst, handler.Filename))
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
