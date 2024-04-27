package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if err := app(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func app() error {
	mux := http.NewServeMux()

	mux.Handle("GET /", viewHandler())
	mux.Handle("POST /", fileUploadHandler())

	addr := ":8080"
	fmt.Printf("Listening on %v...\n", addr)
	return http.ListenAndServe(addr, mux)
}
