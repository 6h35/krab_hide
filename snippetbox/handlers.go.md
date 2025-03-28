package main

import (
	"errors"
	"fmt"

	// "html/template"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"snippetbox.alexedwards.net/internal/models"

	"github.com/julienschmidt/httprouter"
)

type snippetCreateForm struct {
	Title string
	Content string
	Expires int
	FieldErrors map[string]string
	
	}	
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

//	func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
//		title := "O snail"
//		content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
//		expires := 7
//		id, err := app.snippets.Insert(title, content, expires)
//		if err != nil {
//			app.serverError(w, err)
//			return
//		}
//		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
//	}

// func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
// 	// First we call r.ParseForm() which adds any data in POST request bodies
// 	// to the r.PostForm map. This also works in the same way for PUT and PATCH
// 	// requests. If there are any errors, we use our app.ClientError() helper to
// 	// send a 400 Bad Request response to the user.
// 	err := r.ParseForm()
// 	if err != nil {
// 		app.clientError(w, http.StatusBadRequest)
// 		return
// 	}
// 	// Use the r.PostForm.Get() method to retrieve the title and content
// 	// from the r.PostForm map.
// 	title := r.PostForm.Get("title")
// 	content := r.PostForm.Get("content")
// 	// The r.PostForm.Get() method always returns the form data as a *string*.
// 	// However, we're expecting our expires value to be a number, and want to
// 	// represent it in our Go code as an integer. So we need to manually convert
// 	// the form data to an integer using strconv.Atoi(), and we send a 400 Bad
// 	// Request response if the conversion fails.
// 	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
// 	if err != nil {
// 		app.clientError(w, http.StatusBadRequest)
// 		return
// 	}
// 	id, err := app.snippets.Insert(title, content, expires)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
//		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
//	}
// func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		app.clientError(w, http.StatusBadRequest)
// 		return
// 	}
// 	title := r.PostForm.Get("title")
// 	content := r.PostForm.Get("content")
// 	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
// 	if err != nil {
// 		app.clientError(w, http.StatusBadRequest)
// 		return
// 	}
// 	fieldErrors := make(map[string]string)
// 	if strings.TrimSpace(title) == "" {
// 		fieldErrors["title"] = "This field cannot be blank"
// 	} else if utf8.RuneCountInString(title) > 100 {
// 		fieldErrors["title"] = "This field cannot be more than 100 characters long"
// 	}
// 	if strings.TrimSpace(content) == "" {
// 		fieldErrors["content"] = "This field cannot be blank"
// 	}
// 	if expires != 1 && expires != 7 && expires != 365 {
// 		fieldErrors["expires"] = "This field must equal 1, 7 or 365"
// 	}
// 	if len(fieldErrors) > 0 {
// 		fmt.Fprint(w, fieldErrors)
// 		return
// 	}
// 	id, err := app.snippets.Insert(title, content, expires)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
// 	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
// }

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	// Get the expires value from the form as normal.
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	// Create an instance of the snippetCreateForm struct containing the values
	// from the form and an empty map for any validation errors.
	form := snippetCreateForm{
		Title:       r.PostForm.Get("title"),
		Content:     r.PostForm.Get("content"),
		Expires:     expires,
		FieldErrors: map[string]string{},
	}
	// Update the validation checks so that they operate on the snippetCreateForm
	// instance.
	if strings.TrimSpace(form.Title) == "" {
		form.FieldErrors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(form.Title) > 100 {
		form.FieldErrors["title"] = "This field cannot be more than 100 characters long"
	}
	if strings.TrimSpace(form.Content) == "" {
		form.FieldErrors["content"] = "This field cannot be blank"
	}

	if form.Expires != 1 && form.Expires != 7 && form.Expires != 365 {
		form.FieldErrors["expires"] = "This field must equal 1, 7 or 365"
	}
	// If there are any validation errors re-display the create.tmpl template,
	// passing in the snippetCreateForm instance as dynamic data in the Form
	// field.
	if len(form.FieldErrors) > 0 {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.tmpl", data)
		return
	}
	// We also need to update this line to pass the data from the
	// snippetCreateForm instance to our Insert() method.
	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
