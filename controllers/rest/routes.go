package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RunServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m := map[string]string{
			"message": "TodoApp server up and running",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(m)
	})

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ListTasks(w, r)
		case http.MethodPost:
			CreateTask(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetTask(w, r)
		case http.MethodPut:
			CompleteTask(w, r)
		case http.MethodDelete:
			DeleteTask(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Listening on port 8989")
	http.ListenAndServe(":8989", nil)
}
