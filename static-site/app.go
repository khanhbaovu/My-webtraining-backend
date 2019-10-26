package main

import (
//  "fmt"
  "net/http"
  "log"
  "html/template"
  "path/filepath"
)
func templateServer(w http.ResponseWriter, r *http.Request) {
  lp := filepath.Join("templates", "layout.html")
  mp := filepath.Join("templates", filepath.Clean(r.URL.Path))

  tmp := template.Must(template.ParseFiles(lp, mp))

  tmp.ExecuteTemplate(w, "layout", nil)
}
func main() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.HandleFunc("/", templateServer)
  log.Fatal(http.ListenAndServe(":3000", nil))
}
