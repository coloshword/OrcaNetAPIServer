package main

import (
    "fmt"
    "io"
    "net/http"
    "errors"
) 


// some basic endpoints

// getRoot: the root endpoint ('/')
func getRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Println("got / request")
    io.WriteString(w, "This is the root of the API server")
}

// getHello: the hello endpoint
func getHello(w http.ResponseWriter, r *http.Request) {
    fmt.Println("got hello request")
    io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
    http.HandleFunc("/", getRoot)
    http.HandleFunc("/hello", getHello)

    err := http.ListenAndServe(":3333", nil)
    
    if errors.Is(err, http.ErrServerClosed) {
        fmt.Println("server is closed")
    } else if err != nil {
        fmt.Printf("error starting the server %s\n ", err)
    }
}

