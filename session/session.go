package main

import (
    "fmt"
    "net/http"
    "html/template"
    "github.com/gorilla/sessions"
)

var (
  key = []byte("secret-key")
  store = sessions.NewCookieStore(key)
)

func student(w http.ResponseWriter, r *http.Request) {
  session,_ := store.Get(r,"cookie-name")

  if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
    http.Error(w,"You are not Dalhousie Student", http.StatusForbidden)
    return
  }

  tmp,_ := template.ParseFiles("student.html")

  tmp.Execute(w,nil)

}

func login(w http.ResponseWriter, r *http.Request) {
  session,_ := store.Get(r,"cookie-name")

  fmt.Fprintf(w,"Welcome to DalOn")
  session.Values["authenticated"]  = true
  session.Save(r,w)
}

func logout(w http.ResponseWriter, r *http.Request) {
  session,_ := store.Get(r,"cookie-name")

  fmt.Fprintf(w,"Bye Bye")
  session.Values["authenticated"] = false
  session.Save(r,w)
}

func main() {
  http.HandleFunc("/student", student)
  http.HandleFunc("/login", login)
  http.HandleFunc("/logout", logout)

  http.ListenAndServe(":8080", nil)
}
