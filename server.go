// This repo serves as an initial exploration of using GoLang with Negroni.
// The Render library is included to provide diversity in response format support.
// relational data is stored using PostGres via the pq package
// Redis is used as an ultra fast data store for json objects

// More to come...

package main

import (
  "github.com/codegangsta/negroni"
  "log"
  "net/http"
  "database/sql"
  _"github.com/lib/pq"
  "fmt"
  "time"
)


// initialize redis database connection pool
var pool = newPool()

func main() {

  // **************************** start Postgres related code *******************************************
  // create the statement string
  var sStmt string = "insert into users (username, password) values (Eric, Password1)"

  // lazily open db (doesn't truly open until first request)
  db, err := sql.Open("postgres","host=localhost dbname=testdb sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }

  stmt, err := db.Prepare(sStmt)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("StartTime: %v\n", time.Now())

  res, err := stmt.Exec(1, time.Now())
  if err != nil || res == nil {
    log.Fatal(err)
  }

  // close statement
  stmt.Close()

  // close db
  db.Close()

  fmt.Printf("StopTime: %v\n", time.Now())
  // **************************** end Postgres related code ********************************************







  // **************************** start Redis related code *********************************************
  // placeholder database code
  c := pool.Get()
  defer c.Close()
  test,_:=c.Do("HGETALL", "test:1")
  fmt.Println(test)
  // **************************** end Redis related code ***********************************************







  // initialize the routing
  mux := http.NewServeMux()

  // map urls to route handler functions
  mux.HandleFunc("/", HtmlIndex)
  mux.HandleFunc("/json", JsonIndex)


  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":3000")
}
