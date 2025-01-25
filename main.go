package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := Router()

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Server started on http://localhost:8000")
}

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Welcome To This API"))
		if err != nil {
			fmt.Printf("Error writing response: %v\n", err)
		}
	})

	router.HandleFunc("/tasks", HandleTasks).Methods("GET", "POST")
	router.HandleFunc("/tasks/{id}", GetTask)

	return router
}
