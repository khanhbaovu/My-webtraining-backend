package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
)

type student struct {
  Name string
  Major string
  ID string
}

type Students []student

var students Students

func StudentList(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json")

  students = append(students,student{r.FormValue("firstname"),r.FormValue("major"),r.FormValue("id")})

  json.NewEncoder(w).Encode(students)
}

func Student(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json")
  vars := mux.Vars(r)

  name := vars["name"]

  for _,student := range students {
    if student.Name == name {
      json.NewEncoder(w).Encode(student)
    }
  }
}

func AddStudent(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json")
  var newStudent student

  json.NewDecoder(r.Body).Decode(&newStudent)

  students = append(students,newStudent)

  json.NewEncoder(w).Encode(students)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json")
  vars := mux.Vars(r)

  targetName := vars["name"]

  var target student
  
  _ = json.NewDecoder(r.Body).Decode(&target)

  target.Name = targetName

  index := 0

  for _,student := range students {
    if student.Name != targetName {
      index++
    }else {
      break
    }
  }

  temp := students[index+1:]

  students = append(students[:index],target)

  students = append(students,temp...)

  json.NewEncoder(w).Encode(students)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json")
  vars := mux.Vars(r)

  name := vars["name"]

  index := 0

  for _,student := range students {
    if student.Name != name{
      index++
    }else {
      break
    }
  }

  copy(students[index:],students[index+1:])

  students = students[:len(students)-1]

  json.NewEncoder(w).Encode(students)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome to my service")
  fmt.Println("Endpoint Hit: homePage")
}

func RequestHandler() {
  r := mux.NewRouter().StrictSlash(true)

  students = Students{
    student{"Khanh","Computer Science", "833024"},
    student{"Bao", "Software Engineering", "321332"},
    student{"Vu", "Computer Engineering", "757248"}
  }

  r.HandleFunc("/", HomePage)
  r.HandleFunc("/studentlist", StudentList).Methods("GET")
  r.HandleFunc("/studentlist/{name}", Student).Methods("GET")
  r.HandleFunc("/studentlist",AddStudent).Methods("POST")
  r.HandleFunc("/studentlist/{name}", UpdateStudent).Methods("PUT")
  r.HandleFunc("/studentlist/{name}", DeleteStudent).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {

  RequestHandler()

}
