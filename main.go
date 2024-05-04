package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/samherrmann/capture/configuration"
)

func main() {
	if err := app(); err != nil {
		logError(err)
	}
}

func app() error {
	config, err := configuration.Load(os.Args)
	if err != nil {
		return err
	}

	fileUploadHandler, err := newFileUploadHandler(config.Destination)
	if err != nil {
		return err
	}

	viewHandler, err := newViewHandler()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()

	mux.Handle("GET /", viewHandler)
	mux.Handle("POST /", fileUploadHandler)

	fmt.Printf("Listening on %v...\n", config.Address)
	return http.ListenAndServe(config.Address, mux)
}

func logError(err error) {
	// If the help flag was invoked, then print the error to standard output and
	// exit with status 0.
	if configuration.IsHelpError(err) {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}
