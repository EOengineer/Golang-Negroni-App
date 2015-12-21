package main

// Prep commands for DB
// ****************************************************************
// $ initdb ~/Programming/postgres-E utf8
// $ pg_ctl -D ~/Programming/postgres -l ~/logfile start
// $ createdb testdb
// $ psql testdb
// testdb=#

import (
  // "log"
  // "net/http"
  // "database/sql"
  // _"github.com/lib/pq"
)


// func FindByID(resp http.ResponseWriter, req *http.Request) {
//   db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
//   if err != nil {
//     log.Fatal(err)
//   }

//   id := 21
//   rows, err := db.Query("SELECT name FROM users WHERE id = $1", id)
// }


// func FindByID(resp http.ResponseWriter, req *http.Request) {
//   db, err := sql.Open("postgres", "host=localhost dbname=testdb sslmode=disable")
//   if err != nil {
//     log.Fatal(err)
//   }

//   id := 21
//   rows, err := db.Query("SELECT name FROM users WHERE id = $1", id)
// }
