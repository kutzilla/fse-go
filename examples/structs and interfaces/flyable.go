package main

import "fmt"

type Vehicle struct {
  brand string
}

// Defintion eines Interfaces für Flugfahrzeuge
type Flyable interface {
  flyUp(height int)
}

// Implementierung eines Flugzeug
type Plane struct {
  Vehicle
  typ string
}

// Implementierung des Flyable-Interface
func (this *Plane) flyUp(height int) {
  fmt.Println(this.brand, this.typ, "flies",
                height, "m up")
}

func main() {
  var flyable Flyable
  flyable = &Plane{Vehicle{"Airbus"}, "A380"}
  flyable.flyUp(10)
}
