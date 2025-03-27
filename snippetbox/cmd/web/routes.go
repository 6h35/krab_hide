package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// func (app *application) routes() http.Handler { //http.Handler là một interface, không phải struct nên không cần dấu *.
// 	mux := http.NewServeMux()

// 	fileServer := http.FileServer(http.Dir("./ui/static/"))
// 	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

// 	mux.HandleFunc("/", app.home)
// 	mux.HandleFunc("/snippet/view", app.snippetView)
// 	mux.HandleFunc("/snippet/create", app.snippetCreate)

// 	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
// }

// func (app *application) routes() http.Handler {
// 	mux := http.NewServeMux()
// 	fileServer := http.FileServer(http.Dir("./ui/static/"))
// 	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
// 	mux.HandleFunc("/", app.home)
// 	mux.HandleFunc("/snippet/view", app.snippetView)
// 	mux.HandleFunc("/snippet/create", app.snippetCreate)
// 	// Create a middleware chain containing our 'standard' middleware
// 	// which will be used for every request our application receives.
// 	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
// 	// Return the 'standard' middleware chain followed by the servemux.
// 	return standard.Then(mux)
// }

// func (app *application) routes() http.Handler {
// 	// Initialize the router.
// 	router := httprouter.New()
// 	// Update the pattern for the route for the static files.
// 	fileServer := http.FileServer(http.Dir("./ui/static/"))
// 	router.Handler(http.MethodGet, "/static/*filepath",
// 		http.StripPrefix("/static", fileServer))
// 	// And then create the routes using the appropriate methods, patterns and
// 	// handlers.

// 	router.HandlerFunc(http.MethodGet, "/", app.home)
// 	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
// 	router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
// 	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)
// 	// Create the middleware chain as normal.
// 	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
// 	// Wrap the router with the middleware and return it as normal.
// 	return standard.Then(router)
// }

func (app *application) routes() http.Handler {
	router := httprouter.New()
	// Create a handler function which wraps our notFound() helper, and then
	// assign it as the custom handler for 404 Not Found responses. You can also
	// set a custom handler for 405 Method Not Allowed responses by setting
	// router.MethodNotAllowed in the same way too.
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath",
		http.StripPrefix("/static", fileServer))
	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
	router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}
