package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Change the signature of the home handler so it is defined as a method against
// *application.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice
	files := []string{"./ui/html/pages/home.tmpl", "./ui/html/base.tmpl", "./ui/html/partials/nav.tmpl"}
	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	// ----- ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")\

	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we can pass the slice of file
	// paths as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Because the home handler function is now a method against application
		// it can access its fields, including the error logger. We'll write the log
		// message to this instead of the standard logger
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal server error:", 500)
		return
	}

	// We then use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in, which for now we'll
	// leave as nil.
	//----- err = ts.Execute(w, nil)

	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal server error:", 500)
		return
	}
}

// Add a snippetView handler function.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404 page
	// not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a snippetCreate handler function.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	if r.Method != "POST" {
		// Use the Header().Set() method to add an 'Allow: POST' header to the
		// response header map. The first parameter is the header name, and
		// the second parameter is the header value.
		w.Header().Set("Allow", "POST")

		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body. We then return from the function so that the
		// subsequent code is not executed.
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))\

		// Use the http.Error() function to send a 405 status code and "Method Not
		// Allowed" string as the response body.
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet"))
}
