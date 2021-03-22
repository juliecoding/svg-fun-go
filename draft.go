package main

import (
	// "log"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("/Users/juliekohler/go/src/github.com/juliecoding/svg-fun-go"))
	http.Handle("/circle", http.HandlerFunc(circle))
	http.Handle("/gopher.jpg", http.HandlerFunc(showJpg))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	log.Fatal(http.ListenAndServe(":9100", fs))
}