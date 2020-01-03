package main

import (
  "fmt"
	"database/sql"
  "log"
	_ "github.com/lib/pq"
)

func main() {
  connStr := "host=localhost port=5432 user=postgres password=Vbk02122000 dbname=postgres sslmode=disabled"

  db, err := sql.Open("postgres", connStr)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(db)

}
