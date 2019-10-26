package main

import (
  "net/http"
  "html/template"
  "log"
)


/*type ToDo struct {
  Title string
  Done bool
}

type ToDoPage struct {
  PageTitle string
  Todos []ToDo
}*/

func main() {
  tmp := template.Must(template.ParseFiles("layout.html"))
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
/*    data := ToDoPage {
      PageTitle: "My To-Do List",
      Todos : []ToDo {
        {Title:"Becoming a good computer science student", Done: false},
        {Title:"Becoming a software engineer", Done:false},
        {Title:"Starting my own tech start-up", Done:false},
        {Title:"Working Hard", Done:true},
      },
    }*/

    tmp.Execute(w, nil)
  })
  log.Fatal(http.ListenAndServe(":8080", nil))
}
