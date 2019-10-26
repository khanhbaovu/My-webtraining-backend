package main

import (
  "fmt"
  "net/http"
  "log"
  "html/template"
)

func middleware(hd http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    hd(w,r)
    fmt.Fprintf(w,"This is the link of your account: %s", r.URL.Path)
  }
}

func studentAccount(w http.ResponseWriter, r *http.Request) {
  tmp := template.Must(template.ParseFiles("student.html"))
  tmp.Execute(w, nil)
}

func professorAccount(w http.ResponseWriter, r *http.Request) {
  tmp := template.Must(template.ParseFiles("professor.html"))
  tmp.Execute(w, nil)
}

func main() {
  http.HandleFunc("/studentAccount", middleware(studentAccount))
  http.HandleFunc("/professorAccount", middleware(professorAccount))

  log.Fatal(http.ListenAndServe(":8080", nil))
}
