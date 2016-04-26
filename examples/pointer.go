package main

import "fmt"

func main() {
  x := 5
  zero(&x)
  fmt.Println("main: x =", x)
}

func zero(x *int) {
  *x = 0
  fmt.Println("zero: x =", *x)
}
