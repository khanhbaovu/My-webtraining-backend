package main

import (
//  "fmt"
	"database/sql"
  "log"
	_ "github.com/lib/pq"
)
type Tag struct {
    ID   string    `json:"id"`
    Name string `json:"name"`
}
func main() {
  connStr := "host=localhost port=5432 user=postgres password=inininub dbname=postgres sslmode=disable"

  db, err := sql.Open("postgres", connStr)

  if err != nil {
    log.Fatal(err)
  }

  defer db.Close()
  db.Query("INSERT INTO students VALUES ('B00977056', 'Nguyen Huu Hai', 8)")
  results, err := db.Query("SELECT id, name FROM students")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    for results.Next() {
        var tag Tag
        // for each row, scan the result into our tag composite object
        err = results.Scan(&tag.ID, &tag.Name)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        log.Printf(tag.Name)
    }

}
