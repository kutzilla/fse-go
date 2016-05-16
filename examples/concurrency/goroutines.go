package main

import "fmt"

func run(id int) {
  for i := 0; i < 100; i++ {
    fmt.Println(id, "called")
  }
}

func main() {
  for i := 0; i < 10; i++ {
    go run(i + 1)
  }
  var input string
  fmt.Scanln(&input)
}
