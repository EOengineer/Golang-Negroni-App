// This repo serves as an initial exploration of using GoLang with Negroni.
// The Render library is included to provide diversity in response format support.
// relational data is stored using PostGres via the pq package
// Redis is used as an ultra fast data store for json objects

// More to come...

package main

import (
  "github.com/codegangsta/negroni"
  "net/http"
  "fmt"
)


// initialize redis database connection pool
var pool = newPool()

func main() {
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
  mux.HandleFunc("/user", FindByUsername)


  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":3000")
}
