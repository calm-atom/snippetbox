package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slic containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/"
	// If it doesn't send a 404 response to the client
	// We then return from the handler otherwise we would keep executing
	// and also write "Hello from SnippetBox"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// snippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Register two new handler functions for creating and viewing snippets
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Print a log message to say that the server is starting
	log.Print("starting server on :4000")

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on and the servemux we just
	// created. If http.ListenAndServe() returns an error we use the log.Fatal() function
	// to log the error message and exit. Note that any error returned by http.ListenAndServe()
	// is always non-nil.

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
