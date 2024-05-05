package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/samherrmann/capture/cookies"
)

func newFileUploadHandler(dst string) (http.Handler, error) {
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		err := fmt.Errorf("making destination directory: %w", err)
		return nil, err
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusMsg := "Success!"
		statusCode, err := copyFormFile(r, dst)
		if err != nil {
			statusMsg = fmt.Sprintf("Error: %s", err.Error())
		}
		cookies.SetStatus(w, statusCode, statusMsg)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}), nil
}

func copyFormFile(r *http.Request, dst string) (int, error) {
	// maxMemory is set to 32 megabytes.
	if err := r.ParseMultipartForm(32 * 1024 * 1024); err != nil {
		err := fmt.Errorf("parsing multi-part form: %w", err)
		return http.StatusBadRequest, err
	}

	// Get file from form.
	srcFile, handler, err := r.FormFile("file")
	if err != nil {
		err := fmt.Errorf("reading form file: %w", err)
		return http.StatusBadRequest, err
	}
	defer srcFile.Close()

	// Create a filename based on the current Unix time.
	ext := filepath.Ext(handler.Filename)
	filename := fmt.Sprintf("%v%s", time.Now().UnixNano(), ext)

	// Create file on disk.
	dstFile, err := os.Create(filepath.Join(dst, filename))
	if err != nil {
		err := fmt.Errorf("creating destination file: %w", err)
		return http.StatusInternalServerError, err
	}
	defer dstFile.Close()

	// Copy form file into disk file.
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		err := fmt.Errorf("copying source file to destination file: %w", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
