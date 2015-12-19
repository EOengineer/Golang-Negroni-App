// This repo serves as an initial exploration of using GoLang with Negroni.
// The Render library is included to provide diversity in response format support.
// Redis is used as an ultra fast data store for json objects

// More to come...

package main

import (
  "github.com/codegangsta/negroni"
  "net/http"
  "fmt"
  "github.com/unrolled/render"
  "github.com/garyburd/redigo/redis"
)



// set up database pool
func newPool() *redis.Pool {
return &redis.Pool{
          MaxIdle: 80,
          MaxActive: 12000, // max number of connections
          Dial: func() (redis.Conn, error) {
                  c, err := redis.Dial("tcp", ":6379")
                  if err != nil {
                          panic(err.Error())
                  }
                  return c, err
          },
  }

}




// loosely defined JSON struct
type ExampleJSON struct {
  ID      string   `json:"one,attr"`
  Name    string   `json:"two,attr"`
  Gender  string   `json:"three,attr"`
}



// initialize pool
var pool = newPool()

func main() {

  // placeholder database code
  c := pool.Get()
    defer c.Close()
    test,_:=c.Do("HGETALL", "test:1")
    fmt.Println(test)


  // some basic routing
  r := render.New()
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
  })


  // our json response route
  mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
    r.JSON(w, http.StatusOK, ExampleJSON{ID: "1", Name: "Ted", Gender: "Male"})
  })

  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":3000")
}
