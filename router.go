// Copyright © 2018 budougumi0617 All Rights Reserved.

package main

import (
	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// GetTodoRouter returns simple JSON API server
func GetTodoRouter() chi.Router {
	router := chi.NewRouter()
	// Set output for logging.
	middleware.DefaultLogger = middleware.RequestLogger(
		&middleware.DefaultLogFormatter{
			Logger: newLogger(),
		},
	)
	router.Use(middleware.Logger)
	router.HandleFunc("/*", Index) // WildCard.
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoID}", TodoShow)
	router.Post("/todos", TodoCreate)
	return router
}

func newLogger() *log.Logger {
	return log.New(os.Stdout, "chi-log: ", log.Lshortfile)
}
