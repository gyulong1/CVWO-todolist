package main

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	// initializes connectinon to pg database
	db_Init()
	router := mux.NewRouter()

	//router.HandleFunc("/api/login", middleware.Login).Methods("POST", "OPTIONS")
	//router.HandleFunc("/api/register", middleware.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task", GetTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/toggleTask/{id}", ToggleTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", DeleteTask).Methods("DELETE", "OPTIONS")

	return router
}
