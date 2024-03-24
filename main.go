package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Page struct {
	Data string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<p>This is some data loaded via htmx!</p>
	`)
}

func projHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/proj.html")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/data", dataHandler)
	r.HandleFunc("/projects", projHandler)

	// Serve static files from the static/ directory
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.Handle("/", r)

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
