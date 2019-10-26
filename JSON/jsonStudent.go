// json.go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "html/template"
    "log"
)

type formData struct {
  Firstname string
  Lastname string
  Email string
  Major string
  Message string
}

func studentForm(w http.ResponseWriter, r *http.Request) {
  tmp := template.Must(template.ParseFiles("form1.html"))

  if r.Method != http.MethodPost {
    tmp.Execute(w, nil)
    return
  }

  studentData := formData{
    Firstname : r.FormValue("firstname"),
    Lastname : r.FormValue("lastname"),
    Email : r.FormValue("email"),
    Major : r.FormValue("major"),
    Message : r.FormValue("message"),
  }

  fmt.Fprintf(w,"This is your JSON data\n")

  json.NewEncoder(w).Encode(studentData)

}

func main() {
  http.HandleFunc("/", studentForm)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
