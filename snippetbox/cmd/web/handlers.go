package main

import (
	"errors"
	"fmt"

	// "html/template"
	"net/http"
	"strconv"

	"snippetbox.alexedwards.net/internal/models"

	"github.com/julienschmidt/httprouter"
)

// func (app *application) home(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		app.notFound(w) // Use the notFound() helper
// 		return
// 	}
// 	files := []string{
// 		"./ui/html/base.tmpl",
// 		"./ui/html/partials/nav.tmpl",
// 		"./ui/html/pages/home.tmpl",
// 	}
// 	ts, err := template.ParseFiles(files...)
// 	if err != nil {
// 		app.serverError(w, err) // Use the serverError() helper.
// 		return
// 	}
// 	err = ts.ExecuteTemplate(w, "base", nil)
// 	if err != nil {
// 		app.serverError(w, err) // Use the serverError() helper.
// 	}

// }

// func (app *application) home(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		app.notFound(w)
// 		return
// 	}
// 	snippets, err := app.snippets.Latest()
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
// 	for _, snippet := range snippets {
// 		fmt.Fprintf(w, "%+v\n", snippet)
// 	}

// }

// func (app *application) home(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		app.notFound(w)
// 		return
// 	}
// 	snippets, err := app.snippets.Latest()
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
// 	files := []string{
// 		"./ui/html/base.tmpl",
// 		"./ui/html/partials/nav.tmpl",
// 		"./ui/html/pages/home.tmpl",
// 	}
// 	ts, err := template.ParseFiles(files...)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
// 	// Create an instance of a templateData struct holding the slice of
// 	// snippets.

// 	data := &templateData{
// 		Snippets: snippets,
// 	}
// 	// Pass in the templateData struct when executing the template.
// 	err = ts.ExecuteTemplate(w, "base", data)
// 	if err != nil {
// 		app.serverError(w, err)
// 	}
// }

// func (app *application) home(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		app.notFound(w)
// 		return
// 	}
// 	snippets, err := app.snippets.Latest()
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
// 	app.render(w, http.StatusOK, "home.tmpl", &templateData{
// 		Snippets: snippets,
// 	})
// }
// func (app *application) home(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		app.notFound(w)
// 		return
// 	}
// 	snippets, err := app.snippets.Latest()
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
// 	data := app.newTemplateData(r)
// 	data.Snippets = snippets
// 	app.render(w, http.StatusOK, "home.tmpl", data)
// }

// func (app *application) home(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		app.notFound(w)
// 		return
// 	}
// 	snippets, err := app.snippets.Latest()
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
// 	data := app.newTemplateData(r)
// 	data.Snippets = snippets
// 	app.render(w, http.StatusOK, "home.tmpl", data)
// }

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Because httprouter matches the "/" path exactly, we can now remove the
	// manual check of r.URL.Path != "/" from this handler.
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := app.newTemplateData(r)
	data.Snippets = snippets
	app.render(w, http.StatusOK, "home.tmpl", data)
}

// func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil || id < 1 {
// 		app.notFound(w)
// 		return
// 	}
// 	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
// }

// func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil || id < 1 {
// 		app.notFound(w)
// 		return
// 	}
// 	snippet, err := app.snippets.Get(id)
// 	if err != nil {
// 		if errors.Is(err, models.ErrNoRecord) {
// 			app.notFound(w)
// 		} else {
// 			app.serverError(w, err)
// 		}
// 		return
// 	}
// 	fmt.Fprintf(w, "%+v", snippet)
// }

// func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil || id < 1 {
// 		app.notFound(w)
// 		return
// 	}
// 	snippet, err := app.snippets.Get(id)
// 	if err != nil {
// 		if errors.Is(err, models.ErrNoRecord) {
// 			app.notFound(w)
// 		} else {
// 			app.serverError(w, err)
// 		}
// 		return
// 	}
// 	files := []string{
// 		"./ui/html/base.tmpl",
// 		"./ui/html/partials/nav.tmpl",
// 		"./ui/html/pages/view.tmpl",
// 	}
// 	ts, err := template.ParseFiles(files...)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
// 	err = ts.ExecuteTemplate(w, "base", snippet)
// 	if err != nil {
// 		app.serverError(w, err)
// 	}
// }

// func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil || id < 1 {
// 		app.notFound(w)
// 		return
// 	}
// 	snippet, err := app.snippets.Get(id)
// 	if err != nil {
// 		if errors.Is(err, models.ErrNoRecord) {
// 			app.notFound(w)
// 		} else {
// 			app.serverError(w, err)
// 		}
// 		return
// 	}
// 	files := []string{
// 		"./ui/html/base.tmpl",
// 		"./ui/html/partials/nav.tmpl",
// 		"./ui/html/pages/view.tmpl",
// 	}
// 	ts, err := template.ParseFiles(files...)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}

// 	data := &templateData{
// 		Snippet: snippet,
// 	}

// 	err = ts.ExecuteTemplate(w, "base", data)
// 	if err != nil {
// 		app.serverError(w, err)
// 	}
// }

// func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil || id < 1 {
// 		app.notFound(w)
// 		return
// 	}
// 	snippet, err := app.snippets.Get(id)
// 	if err != nil {
// 		if errors.Is(err, models.ErrNoRecord) {
// 			app.notFound(w)
// 		} else {
// 			app.serverError(w, err)
// 		}
// 		return
// 	}
// 	// Use the new render helper.
// 	app.render(w, http.StatusOK, "view.tmpl", &templateData{
// 		Snippet: snippet,
// 	})
// }

// func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil || id < 1 {
// 		app.notFound(w)
// 		return
// 	}
// 	snippet, err := app.snippets.Get(id)
// 	if err != nil {
// 		if errors.Is(err, models.ErrNoRecord) {
// 			app.notFound(w)
// 		} else {
// 			app.serverError(w, err)
// 		}
// 		return
// 	}
// 	data := app.newTemplateData(r)
// 	data.Snippet = snippet
// 	app.render(w, http.StatusOK, "view.tmpl", data)
// }

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// When httprouter is parsing a request, the values of any named parameters
	// will be stored in the request context. We'll talk about request context
	// in detail later in the book, but for now it's enough to know that you can
	// use the ParamsFromContext() function to retrieve a slice containing these
	// parameter names and values like so:
	params := httprouter.ParamsFromContext(r.Context())
	// We can then use the ByName() method to get the value of the "id" named
	// parameter from the slice and validate it as normal.
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	data := app.newTemplateData(r)

	data.Snippet = snippet
	app.render(w, http.StatusOK, "view.tmpl", data)
}

// func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		w.Header().Set("Allow", http.MethodPost)
// 		app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper.
// 		return
// 	}
// 	w.Write([]byte("Create a new snippet..."))
// }

// func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		w.Header().Set("Allow", http.MethodPost)
// 		app.clientError(w, http.StatusMethodNotAllowed)
// 		return
// 	}
// 	// Create some variables holding dummy data. We'll remove these later on
// 	// during the build.
// 	title := "O snail"
// 	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
// 	expires := 7
// 	// Pass the data to the SnippetModel.Insert() method, receiving the
// 	// ID of the new record back.
// 	id, err := app.snippets.Insert(title, content, expires)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
// 	// Redirect the user to the relevant page for the snippet.
// 	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id),
// 		http.StatusSeeOther)
// }

// func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Display the form for creating a new snippet..."))
// }

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "create.tmpl", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Checking if the request method is a POST is now superfluous and can be
	// removed, because this is done automatically by httprouter.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Update the redirect path to use the new clean URL format.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
