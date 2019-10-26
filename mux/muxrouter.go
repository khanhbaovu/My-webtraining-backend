package main

import (
  "fmt"
  "net/http"
  "log"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/book/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    title := vars["title"]
    page := vars["page"]

    fmt.Fprintf(w,"This book is %s and the page is %s!", title, page)
  })

  log.Fatal(http.ListenAndServe(":8080", r))
}
