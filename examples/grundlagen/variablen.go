package main
import (
  "fmt"
  "os"
)
const greet ="Hello, "

func main() {
  var name string
  name = os.Args[1]

  greeting := greet + name
  fmt.Println(greeting)
}
