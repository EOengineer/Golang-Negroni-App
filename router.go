package main

import (
  "fmt"
  "net/http"
  "github.com/unrolled/render"
)



// our HTML response
func HtmlIndex(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "Welcome to the home page!")
}

// our JSON response
func JsonIndex(w http.ResponseWriter, req *http.Request) {
  r := render.New()
  r.JSON(w, http.StatusOK, ExampleJSON{ID: "1", Name: "Ted", Gender: "Male"})
}

