package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Println("Hi!")
  fmt.Println(runtime.GOMAXPROCS(-1))
}
