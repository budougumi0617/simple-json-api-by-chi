// Copyright © 2018 budougumi0617 All Rights Reserved.

package main

import (
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
	fmt.Fprintln(w, "Todo Index!")
}

// TodoShow is not implemented.
func TodoShow(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID") // Get parameter form URL PATH.
	fmt.Fprintln(w, "Todo show:", todoID)
}
