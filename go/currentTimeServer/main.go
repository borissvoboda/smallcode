package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
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

	// serve static html file
	http.HandleFunc("/in", func(w http.ResponseWriter, r *http.Request) {
        html, _ := os.ReadFile("index.html")
        w.Header().Set("Content-Type", "text/html")
        w.Write(html)
    })

	var tmpl = template.Must(template.ParseGlob("inj.html"))

	http.HandleFunc("/inj", func(w http.ResponseWriter, r *http.Request) {
    data := struct{ Name string }{Name: "My name"}  
    tmpl.ExecuteTemplate(w, "inj.html", data)
})


	http.ListenAndServe(":8080", nil)

}

