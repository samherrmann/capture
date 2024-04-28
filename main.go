package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/samherrmann/capture/configuration"
)

func main() {
	if err := app(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func app() error {
	config, err := configuration.Load(os.Args)
	if err != nil {
		return fmt.Errorf("loading configuration: %w", err)
	}

	fileUploadHandler, err := newFileUploadHandler(config.Destination)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()

	mux.Handle("GET /", viewHandler())
	mux.Handle("POST /", fileUploadHandler)

	fmt.Printf("Listening on %v...\n", config.Address)
	return http.ListenAndServe(config.Address, mux)
}
