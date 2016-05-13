package main

import "fmt"

type Vehicle struct {
  brand string
}

func (this *Vehicle) move() {
  fmt.Println(this.brand, "moves")
}

type Car struct {
  Vehicle
  model string
}

func main() {
  var v Vehicle
  v = Car{Vehicle{"Audi"}, "R8"}
}
