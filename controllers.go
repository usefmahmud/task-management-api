package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

var Tasks []Task

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := Response{
		Status: 200,
		Data: func() interface{} {
			if len(Tasks) == 0 {
				return []Task{}
			}
			return Tasks
		}(),
	}
	json.NewEncoder(w).Encode(res)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	text := r.FormValue("text")
	if text == "" {
		res := Response{
			Status: 400,
			Data: map[string]string{
				"error": "text is required",
			},
		}
		json.NewEncoder(w).Encode(res)

		return
	}

	id := addTask(text, &Tasks)
	res := Response{
		Status: 201,
		Data: map[string]string{
			"id": id.String(),
		},
	}

	json.NewEncoder(w).Encode(res)
}

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetTasks(w, r)
	} else if r.Method == http.MethodPost {
		CreateTask(w, r)
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// task_id := vars["id"]

}
