package main

import "fmt"

// Struktur für Fahrzeuge
type Vehicle struct {
  // Definition einer Instanzvariable
  brand string
}

// Defintion einer Methode
func (this *Vehicle) move() {
  fmt.Println(this.brand, "moves")
}

func main() {
  // Deklaration einer Struct-Instanz
  var v Vehicle
  // Initialisierung durch Angabe des Variablennamens
  v = Vehicle{brand: "Audi"}
  v.move()
  // Initialisierung durch Angabe in der Reihenfolge
  v = Vehicle{"VW"}
  v.move()
}
