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
r = information
