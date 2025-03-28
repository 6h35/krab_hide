package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.alexedwards.net/internal/models"

	"github.com/julienschmidt/httprouter"
	"snippetbox.alexedwards.net/internal/validator"
)

type snippetCreateForm struct {
	Title       string
	Content     string
	Expires     int
	FieldErrors map[string]string
	validator.Validator
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
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
// 	params := httprouter.ParamsFromContext(r.Context())
// 	id, err := strconv.Atoi(params.ByName("id"))
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

// func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
// 	params := httprouter.ParamsFromContext(r.Context())
// 	id, err := strconv.Atoi(params.ByName("id"))
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
// 	// Use the PopString() method to retrieve the value for the "flash" key.
// 	// PopString() also deletes the key and value from the session data, so it
// 	// acts like a one-time fetch. If there is no matching key in the session
// 	// data this will return the empty string.
// 	flash := app.sessionManager.PopString(r.Context(), "flash")
// 	data := app.newTemplateData(r)
// 	data.Snippet = snippet
// 	// Pass the flash message to the template.
// 	data.Flash = flash
// 	app.render(w, http.StatusOK, "view.tmpl", data)
// }

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
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
// 	data := app.newTemplateData(r)
// 	app.render(w, http.StatusOK, "create.tmpl", data)
// }

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = snippetCreateForm{
		Expires: 365,
	}
	app.render(w, http.StatusOK, "create.tmpl", data)
}

//	func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
//		err := r.ParseForm()
//		if err != nil {
//			app.clientError(w, http.StatusBadRequest)
//			return
//		}
//		expires, err := strconv.Atoi(r.PostForm.Get("expires"))
//		if err != nil {
//			app.clientError(w, http.StatusBadRequest)
//			return
//		}
//		form := snippetCreateForm{
//			Title:       r.PostForm.Get("title"),
//			Content:     r.PostForm.Get("content"),
//			Expires:     expires,
//			FieldErrors: map[string]string{},
//		}
//		if strings.TrimSpace(form.Title) == "" {
//			form.FieldErrors["title"] = "This field cannot be blank"
//		} else if utf8.RuneCountInString(form.Title) > 100 {
//			form.FieldErrors["title"] = "This field cannot be more than 100 characters long"
//		}
//		if strings.TrimSpace(form.Content) == "" {
//			form.FieldErrors["content"] = "This field cannot be blank"
//		}
//		if form.Expires != 1 && form.Expires != 7 && form.Expires != 365 {
//			form.FieldErrors["expires"] = "This field must equal 1, 7 or 365"
//		}
//		if len(form.FieldErrors) > 0 {
//			data := app.newTemplateData(r)
//			data.Form = form
//			app.render(w, http.StatusUnprocessableEntity, "create.tmpl", data)
//			return
//		}
//		id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
//		if err != nil {
//			app.serverError(w, err)
//			return
//		}
//		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
//	}

// func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		app.clientError(w, http.StatusBadRequest)
// 		return
// 	}
// 	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
// 	if err != nil {
// 		app.clientError(w, http.StatusBadRequest)
// 		return
// 	}
// 	form := snippetCreateForm{
// 		Title:   r.PostForm.Get("title"),
// 		Content: r.PostForm.Get("content"),
// 		Expires: expires,
// 	}
// 	// Perform validation checks using the embedded Validator type.
// 	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
// 	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
// 	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
// 	form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

// 	// If validation fails, re-render the form with error messages.
// 	if !form.Valid() {
// 		data := app.newTemplateData(r)
// 		data.Form = form
// 		app.render(w, http.StatusUnprocessableEntity, "create.tmpl", data)
// 		return
// 	}

// 	// Insert the snippet into the database.
// 	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}

//		// Redirect to the snippet view page.
//		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
//	}

// func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		app.clientError(w, http.StatusBadRequest)
// 		return
// 	}

// 	// Declare a new empty instance of the snippetCreateForm struct.
// 	var form snippetCreateForm

// 	// Call the Decode() method of the form decoder, passing in the current
// 	// request and *a pointer* to our snippetCreateForm struct. This will
// 	// essentially fill our struct with the relevant values from the HTML form.
// 	// If there is a problem, we return a 400 Bad Request response to the client.
// 	err = app.formDecoder.Decode(&form, r.PostForm)
// 	if err != nil {
// 		app.clientError(w, http.StatusBadRequest)
// 		return
// 	}

// 	// Then validate and use the data as normal...
// 	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
// 	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
// 	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
// 	form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

// 	if !form.Valid() {
// 		data := app.newTemplateData(r)
// 		data.Form = form
// 		app.render(w, http.StatusUnprocessableEntity, "create.tmpl", data)
// 		return
// 	}

// 	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}

// 	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
// }

// func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
// 	var form snippetCreateForm
// 	err := app.decodePostForm(r, &form)
// 	if err != nil {
// 		app.clientError(w, http.StatusBadRequest)
// 		return
// 	}
// 	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
// 	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
// 	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
// 	form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")
// 	if !form.Valid() {
// 		data := app.newTemplateData(r)
// 		data.Form = form
// 		app.render(w, http.StatusUnprocessableEntity, "create.tmpl", data)
// 		return
// 	}
// 	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
// 	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
// }

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	var form snippetCreateForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.tmpl", data)
		return
	}
	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Use the Put() method to add a string value ("Snippet successfully created!") and the corresponding key ("flash") to the session data.
	app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
