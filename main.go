// Copyright © 2018 budougumi0617 All Rights Reserved.

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/docgen"
)

func main() {
	router := chi.NewRouter()
	router.Get("/*", Index) // *つけないと全てのPATHを受けない。
	// ルーティング情報を自動生成
	fmt.Println(docgen.MarkdownRoutesDoc(router, newMarkdownOpts()))
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index writes simple response
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
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
