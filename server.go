package main

import (
    "crypto/md5"
    "fmt"
    "log"
    "net/http"
    _ "net/http/pprof"
   "sync"
   "strconv"
)

func work(id int,  wg *sync.WaitGroup) {
  fmt.Println("Starting: ", id)
  for i := 0; i < 1000000; i++ {
    _ = md5.Sum([]byte(strconv.Itoa(i)))
  }
  fmt.Println("Ending: ", id)
  wg.Done()
}

func handler(w http.ResponseWriter, r *http.Request) {
  wg := sync.WaitGroup{}
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go work(i, &wg)
  }
  wg.Wait()
}

func main() {
    http.HandleFunc("/work", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
