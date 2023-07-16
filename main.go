package main

import (
  "fmt"
  "log"
  "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/hello" {
    http.Error(w, "404 not found", http.StatusNotFound)
    return
  }

  if r.Method != "GET" {
    http.Error(w, "Method is not supported", http.StatusNotFound)
    return
  }

  fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
  if err := r.ParseForm(); err != nil {
    fmt.Fprintf(w, "ParseForm() err: %v", err)
    return
  }
  fmt.Fprintf(w, "POST request successful")
  
  name := r.FormValue("name")
  address := r.FormValue("address")

  fmt.Fprintf(w, "Name = %s\n", name)
  fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
  // this will tell Go to search index.html (default) in static dir
  fileServer := http.FileServer(http.Dir("./static))
  http.Handle("/", fileServer)   // root route
  http.HandleFunc("/form", formHandler)  // form route
  http.HandleFunc("/hello", helloHandler)  // hello route

  fmt.Printf("Starting server at port 8000\n")

  // to create server
  if err := http.ListenAndServe(":8000", nil); err !nil {
    log.Fatal(err)
  }
}