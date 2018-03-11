// Copyright © 2018 budougumi0617 All Rights Reserved.

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/docgen"
)

func main() {
	router := chi.NewRouter()
	router.HandleFunc("/*", Index) // WildCard.
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoID}", TodoShow)

	// Auto generate mapping information by markdown format.
	fmt.Println(docgen.MarkdownRoutesDoc(router, newMarkdownOpts()))
	log.Fatal(http.ListenAndServe(":8080", router))
}

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

func newMarkdownOpts() docgen.MarkdownOpts {
	return docgen.MarkdownOpts{
		ProjectPath:        "github.com/budougumi0617/simple-json-api-by-chi",
		Intro:              "Sample JSON API server by go-chi.",
		ForceRelativeLinks: true,
		URLMap: map[string]string{
			"github.com/my/package/vendor/go-chi/chi/": "https://github.com/go-chi/chi/blob/master/",
		},
	}
}
