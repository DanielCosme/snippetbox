package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	r.Use(app.recoverPanic)
	r.Use(app.logRequest)
	r.Use(secureHeaders)

	r.Group(func(r chi.Router) {
		r.Use(app.session.Enable)

		r.Get("/", app.home)
		r.Get("/snippet/create", app.createSnippetForm)
		r.Post("/snippet/create", app.createSnippet)
		r.Get("/snippet/{id}", app.showSnippet)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// Return the 'standard' middleware chain followed by the servemux.
	return r
}
