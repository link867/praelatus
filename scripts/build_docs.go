package main

import (
	"fmt"
	"io/ioutil"

	"github.com/praelatus/backend/api"
	"github.com/pressly/chi/docgen"
)

func main() {
	router := api.New(nil)
	docs := docgen.MarkdownRoutesDoc(router, docgen.MarkdownOpts{
		ProjectPath: "github.com/praelatus/backend",
		Intro:       "This is the REST API reference documentation for the Praelatus REST API, Praelatus is an Open Source Bug Tracking / Ticketing system.",
	})
	err := ioutil.WriteFile("API_REFERENCE.md", []byte(docs), 0744)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Docs generated.")
}
