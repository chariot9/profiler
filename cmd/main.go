package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"profiler/internal/controllers"
	"profiler/internal/middleware"
)

func main() {
	router := mux.NewRouter()

	// List of API endpoints
	router.HandleFunc("/api/users/new", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/info/new", controllers.CreateInfo).Methods("POST")
	router.HandleFunc("/api/me/info", controllers.GetMyInfo).Methods("GET")

	// Use middleware
	router.Use(middleware.JwtAuthentication)

	// Deploy with specific port
	port := os.Getenv("deployment_port")
	if port == "" {
		port = "8000"
	}

	// Start up server
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
