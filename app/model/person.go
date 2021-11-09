package model

type person struct {
	ID        int    `json:"ID,omitempty"`
	FirstName string `json:"FirstName,omitempty"`
	Lastname  string `json:"Lastname,omitempty"`
}

type allTasks []person
