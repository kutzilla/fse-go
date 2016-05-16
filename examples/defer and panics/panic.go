package main

import "fmt"

func divide(dividend float64, divisor float64) float64 {
  if (divisor == 0) {
    //panic("Divisor cannot be zero")
  }
  return dividend / divisor
}

func main() {
  divide(100.0, 0.0)
  defer func() {
    str := recover()
    fmt.Println(str)
  }()
}
