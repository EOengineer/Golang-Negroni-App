package main

import (
  "fmt"
  "net/http"
  "github.com/unrolled/render"
  "log"
  "database/sql"
  _"github.com/lib/pq"
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

// user lookup
func FindByUsername(resp http.ResponseWriter, req *http.Request) {
  db, err := sql.Open("postgres", "host=localhost dbname=Golang_dev sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }

  // var username string = req.FormValue['username']
  var username string = "Eric"
  rows, err := db.Query("SELECT * FROM users WHERE username = $1", username)
  if err != nil {
    log.Fatal(err)
  }
  

  fmt.Printf("User: %s\n", rows)
  // close db
  db.Close()
}
