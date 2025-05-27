package rest

import (
	"TodoApp/services"
	"TodoApp/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//CreateTask

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title string `json:"title"`
		Done  bool   `json:"done"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if req.Title == "" {
		handleError(w, http.StatusBadRequest, "Title is required")
		return
	}

	task, err := services.CreateTask(req.Title, req.Done)
	if err != nil {
		handleError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating task: %v", err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// GetTask by ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ExtractID(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, err.Error())
		return
	}

	task, err := services.GetTask(id)
	if err != nil {
		handleError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving task: %v", err))
		return
	}

	if task == nil {
		handleError(w, http.StatusNotFound, "Task not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// DeleteTask by ID
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ExtractID(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := services.DeleteTask(id); err != nil {
		handleError(w, http.StatusInternalServerError, fmt.Sprintf("Error deleting task: %v", err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// CompleteTask by ID
func CompleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ExtractID(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, err.Error())
		return
	}

	task, err := services.CompleteTask(id)
	if err != nil {
		handleError(w, http.StatusInternalServerError, fmt.Sprintf("Error completing task: %v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// ListTasks get all tasks
func ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks := services.ListTasks()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
