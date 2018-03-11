package main

import "time"

// Todo is a basic model.
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos is Todo list.
type Todos []Todo
