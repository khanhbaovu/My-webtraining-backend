package main

import (
  "fmt"
  "golang.org/x/crypto/bcrypt"
)

func PasswordHasher(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password),14)

  return string(bytes), err
}

func PasswordChecker(password, hash string) bool  {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

  return err == nil
}
func main() {
  password := "khanhbaovu"

  hash,_ := PasswordHasher(password)

  fmt.Println(password)
  fmt.Println(hash)

  check := PasswordChecker(password, hash)

  fmt.Println(check)
}
