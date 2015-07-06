package main

import "time"

//Todo definition
type Todo struct {
	ID        int       `json:"ID"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

//Todos list of Todo
type Todos []Todo
