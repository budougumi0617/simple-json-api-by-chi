// Copyright © 2018 budougumi0617 All Rights Reserved.

package main

import (
	"github.com/go-chi/chi"
)

// GetTodoRouter returns simple JSON API server
func GetTodoRouter() chi.Router {
	router := chi.NewRouter()
	router.HandleFunc("/*", Index) // WildCard.
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoID}", TodoShow)
	return router
}
