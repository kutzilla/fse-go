package main

import "fmt"

type Vehicle struct {
  brand string
}

func (this *Vehicle) move() {
  fmt.Println(this.brand, "moves")
}

type Car struct {
  // Angabe des Embedded Type
  Vehicle
  model string
}

// Überschreiben der move-Methode aus Vehicle
func (this *Car) move() {
  fmt.Println(this.brand, this.model, "moves")
}

func main() {
  // Erzeugung einer Vehicle-Instanz
  audi := Vehicle{brand: "Audi"}
  audi.move()
  // Übergabe des Embedded Type als Variable
  r8 := Car{audi, "R8"}
  // Aufruf der move-Methode der Vehicle-Instanz
  r8.Vehicle.move()
  // Aufruf der move-Methode der Car-Instanz
  r8.move()
  // Direkte Übergabe des Embedded Type
  golf := Car{Vehicle{"VW"}, "Golf"}
  golf.move()
}
