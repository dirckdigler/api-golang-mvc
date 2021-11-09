//Package main provides API to generate a simple CRUD.
// @title User API documentation
// @version 1.0.0

// @host localhost:3000
// @BasePath /tasks
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Persistence.
var tasks = allTasks{
	{
		ID:        1,
		FirstName: "Task One",
		Lastname:  "Some Content",
	},
}

//getTasks get all tasks.
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// getTask get specific task depending ID parameter.
func getTask(w http.ResponseWriter, r *http.Request) {
	// Get parameter.
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, value := range tasks {
		if value.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(value)
		}
	}

}

// deleteTask delete specific task depending ID parameter.
func deleteTask(w http.ResponseWriter, r *http.Request) {
	// Get parameter.
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for i, value := range tasks {
		if value.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "The task ID %v was removed succesfully", taskID)
		}
	}
}

// createTask create task.
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask person
	// ioutil manage enter and output server.
	// Get request information.
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Soemthing went wrong...Insert valid person data")
	}
	// Assing this information to newTask variable.
	json.Unmarshal(reqBody, &newTask)
	// Create an ID dinammically.
	newTask.ID = len(tasks) + 1
	// Save this new information in list of tasks.
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// Respond client with the information.
	json.NewEncoder(w).Encode(newTask)

}

// updateTask update specific task depending ID parameter.
func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask person

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")
	}
	json.Unmarshal(reqBody, &updatedTask)

	for i, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updatedTask.ID = t.ID
			tasks = append(tasks, updatedTask)
			fmt.Fprintf(w, "The task with ID %v has been updated successfully", taskID)
		}
	}

}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/task/{id}", getTask).Methods("GET")
	router.HandleFunc("/delete-task/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/create-task", createTask).Methods("POST")
	router.HandleFunc("/update-task/{id}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}
