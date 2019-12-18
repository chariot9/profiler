package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	port := os.Getenv("deployment_port")

	if port == "" {
		port = "8000"
	}

	_ = http.ListenAndServe(":"+port, router)
}
