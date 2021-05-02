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
		r.Use(noSurf)

		r.Group(func(r chi.Router) {
			r.Use(app.requireAuthentication)

			r.Get("/snippet/create", app.createSnippetForm)
			r.Post("/snippet/create", app.createSnippet)
			r.Post("/user/logout", app.logoutUser)
		})

		r.Get("/", app.home)
		r.Get("/snippet/{id}", app.showSnippet)

		// Add the five new routes.
		r.Get("/user/signup", app.signupUserForm)
		r.Post("/user/signup", app.signupUser)
		r.Get("/user/login", app.loginUserForm)
		r.Post("/user/login", app.loginUser)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// Return the 'standard' middleware chain followed by the servemux.
	return r
}
