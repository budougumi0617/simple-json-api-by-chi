// Copyright © 2018 budougumi0617 All Rights Reserved.

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/docgen"
)

func main() {
	router := GetTodoRouter()

	// Auto generate mapping information by markdown format.
	fmt.Println(docgen.MarkdownRoutesDoc(router, newMarkdownOpts()))
	log.Fatal(http.ListenAndServe(":8080", router))
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
