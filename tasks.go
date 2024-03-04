package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type TasksService struct {
	store Store
}

func NewTasksService(s Store) *TasksService {
	return &TasksService{store: s}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", s.createTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.getTask).Methods("GET")
}

func (s *TasksService) createTask(w http.ResponseWriter, r *http.Request) {
}

func (s *TasksService) getTask(w http.ResponseWriter, r *http.Request) {
}
