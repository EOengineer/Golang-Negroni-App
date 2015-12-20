// This repo serves as an initial exploration of using GoLang with Negroni.
// The Render library is included to provide diversity in response format support.
// Redis is used as an ultra fast data store for json objects

// More to come...

package main

import (
  "github.com/codegangsta/negroni"
  "net/http"
  "fmt"

)


// initialize database connection pool
var pool = newPool()

func main() {

  // placeholder database code
  c := pool.Get()
  defer c.Close()
  test,_:=c.Do("HGETALL", "test:1")
  fmt.Println(test)


  // initialize the routing
  mux := http.NewServeMux()

  // map urls to handler functions
  mux.HandleFunc("/", HtmlIndex)
  mux.HandleFunc("/json", JsonIndex)


  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":3000")
}
