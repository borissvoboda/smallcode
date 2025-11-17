Current Time Server

# package main

Every Go program starts with this line. It means that it is an executable program.

# import (...)

Importing built-in packages

## fmt

For formatted output; in go you need this even for the basic printf!

## net/http

A package needed for the HTTP server

## time

For getting a current time

# main()

Entry point for every program.

## http.HandleFunc("/", func(w http.ResponseWriter, r \*http.Request) {

This is the function that is executed whenever the localhost:8080 is visited.
w = where we write the HTTP response
r = information about the incoming request

## fmt.Fprintf(w, "Current time: %s", time.Now().Format(time.RFC1123))

time.Now() → gets the exact current date/time
.Format(time.RFC1123) → converts it to a standard readable format
Example: Mon, 16 Nov 2025 21:32:45 CET

## http.ListenAndServe(":8080", nil)

Start the web server on port 8080
":8080" → listen on all network interfaces (localhost + network)
nil → use the default router (the one we configured above)
This line blocks forever – server keeps running until you press Ctrl+C

///

# os

# html template

https://pkg.go.dev/html/template
