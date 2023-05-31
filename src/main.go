package main

import (
	"net/http"
)

func main() {

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../pages/assets"))))

	http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("../pages/vendor"))))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../pages/"+r.URL.Path+".html")
	}))

	// Build the Server

	http.ListenAndServe(":8080", nil)
}
