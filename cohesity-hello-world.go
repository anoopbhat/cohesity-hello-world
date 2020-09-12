package main

import (
    "log"
    "net/http"
)

func main() {

    // how we serve up static pages like the forms
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)
    // log that we're going to start listening for requests
    log.Println("Listening...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
