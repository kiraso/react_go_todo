package router

import (
	"github.com/kiraso/react_go_todo/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET","OPTIONS")
	router.HandleFunc("/api/tasks",middleware.CreateTask).Methods("POST","OPTIONS")
	router.HandleFunc("/api/task/{id}",middleware.TaskComplete).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/undoTask/{id}",middleware.UndoTask).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/deleateTask/{id}",middleware.DeleteTask).Methods("DELETE","OPTIONS")
	router.HandleFunc("/api/deleteAllTasks",middleware.DeleteAllTasks).Methods("DELETE","OPTIONS")
	return router
}