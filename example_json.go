package main

// loosely defined JSON struct
type ExampleJSON struct {
  ID      string   `json:"one,attr"`
  Name    string   `json:"two,attr"`
  Gender  string   `json:"three,attr"`
}


type ExampleJSONS []ExampleJSON
