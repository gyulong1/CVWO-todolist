package main

import (
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func router() *mux.Router {

	Init()
	router := mux.NewRouter()
	//router.HandleFunc("/api/login", middleware.Login).Methods("POST", "OPTIONS")
	//router.HandleFunc("/api/register", middleware.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task", GetTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/toggleTask/{id}", ToggleTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", DeleteTask).Methods("DELETE", "OPTIONS")
	return router
}