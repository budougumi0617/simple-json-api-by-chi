// Copyright © 2018 budougumi0617 All Rights Reserved.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// Index returns simple response.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// TodoIndex is not implemented.
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoShow is not implemented.
func TodoShow(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID") // Get parameter form URL PATH.
	fmt.Fprintln(w, "Todo show:", todoID)
}
