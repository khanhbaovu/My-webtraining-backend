package main

import (
    "log"
    "net/http"
    "time"
    "html/template"
)

type MiddleWares func(http.HandlerFunc) http.HandlerFunc

func logger() MiddleWares {
  return func(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      start := time.Now()
      defer func(){
        log.Println(r.URL.Path, time.Since(start))
        }()

      f(w,r)
    }
  }
}

func checker(string string) MiddleWares {
  return func(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      if string == "student" {
        log.Println("StudentAccount")
      }else if string == "professor" {
        log.Println("professor")
      }
      f(w,r)
    }
  }
}

func studentAccount(w http.ResponseWriter, r *http.Request) {
  tmp := template.Must(template.ParseFiles("student.html"))
  tmp.Execute(w,nil)
}

func professorAccount(w http.ResponseWriter, r *http.Request) {
  tmp := template.Must(template.ParseFiles("professor.html"))
  tmp.Execute(w, nil)
}

func middlewareChain(f http.HandlerFunc, middlewares ...MiddleWares) http.HandlerFunc{
  for _,m := range middlewares {
    f = m(f)
  }
  return f
}

func main() {
  http.HandleFunc("/student", middlewareChain(studentAccount,checker("student"),logger()))
  http.HandleFunc("/professor", middlewareChain(professorAccount,checker("professor"),logger()))

  log.Fatal(http.ListenAndServe(":8080", nil))
}
