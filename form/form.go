package main

import (
  "html/template"
  "net/http"
  "log"
)

type formData struct {
  firstname string
  lastname string
  email string
  major string
  message string
}

type messageData struct {
  Success bool
  Firstname string
  Email string
  Major string
}

func formServer(w http.ResponseWriter, r *http.Request) {
  tmp := template.Must(template.ParseFiles("form1.html"))
  if r.Method != http.MethodPost {
    tmp.Execute(w, nil)
    return
  }

  data := formData {
    firstname : r.FormValue("firstname"),
    lastname : r.FormValue("lastname"),
    email : r.FormValue("email"),
    major : r.FormValue("major"),
    message : r.FormValue("message"),
  }

  _ = data
  mdata := messageData {
    Success: true,
    Firstname: data.firstname,
    Email : data.email,
    Major : data.major,
  }
  tmp.Execute(w,mdata)
}
func main() {
  http.HandleFunc("/", formServer)

  log.Fatal(http.ListenAndServe(":8000", nil))
}
