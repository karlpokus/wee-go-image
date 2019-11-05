package main

import (
  "net/http"
  "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "hello!")
}

func main() {
  http.HandleFunc("/hi", handler)
  http.ListenAndServe(":8989", nil)
}
