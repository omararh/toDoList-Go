package router

import (
	"toDoList-go/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET")
	router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE")
	return router
}