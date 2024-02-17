package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", ":8080", "address to serve")
)

func main() {
	http.HandleFunc("/", redirectToFarcasterArtyom)
	http.HandleFunc("/artiom", redirectToFarcasterArtyom)
	http.HandleFunc("/c/surreal", redirectToFarcasterSurreal)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func redirectToFarcasterArtyom(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://warpcast.com/artiom", http.StatusPermanentRedirect)
}

func redirectToFarcasterSurreal(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://warpcast.com/~/channel/surreal", http.StatusPermanentRedirect)
}
