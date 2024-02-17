package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", redirectToFarcaster)
	r.HandleFunc("/{name}", redirectToFarcasterName)
	r.HandleFunc("/c/surreal", redirectToFarcasterSurreal)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err.Error())
	}
}
func redirectToFarcaster(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, " https://warpcast.com", http.StatusSeeOther)
}

func redirectToFarcasterName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	http.Redirect(w, r, "https://warpcast.com/"+name, http.StatusSeeOther)
}

func redirectToFarcasterSurreal(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://warpcast.com/~/channel/surreal", http.StatusSeeOther)
}
