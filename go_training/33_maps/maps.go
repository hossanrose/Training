package main

import "fmt"

type Vertex struct{
  Lat,Long float64
}

var m map[string]Vertex

func main(){
  m = make(map[string]Vertex)
  m["Bell Labs"] = Vertex{
    3434.2323,
    532.3223,
  }
  fmt.Println(m["Bell Labs"])
}
