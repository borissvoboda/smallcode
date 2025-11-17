package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	message := "..."

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Current time: %s %s\n", time.Now().Format(time.RFC1123), message)
		fmt.Fprintf(w, "Current time: %s", time.Now().Format(time.RFC1123))
	})

	// Seems that you can do more of these functions and it works
	http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "a")
	})

	// I want a post now.
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            fmt.Fprintf(w, "POST received! ")
        } else {
            // GET or any other method
            fmt.Fprintf(w, "GET")
        }
    })


	http.ListenAndServe(":8080", nil)

}

