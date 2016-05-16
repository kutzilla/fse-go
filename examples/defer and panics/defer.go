package main

import "fmt"

func main() {
  defer fmt.Println("First")
  fmt.Println("Second")
}
