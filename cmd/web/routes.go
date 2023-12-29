package main

import (
	"net/http"

	"github.com/justinas/alice"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our application receives
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	// Wrap the existing chain with the logRequest middleware.
	return standard.Then(mux)
}
