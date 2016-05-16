package main

import "fmt"

func main() {
	//1. Erstellen eines Slices
	var w []float64

	//Länge 5 mit Kapazität 5 (Array Länge)
	x := make([]float64, 5)

	//Länge 5 Kapazit 10
	y := make([]float64, 5, 10)

	//Slice mit den Positionen 2 bis 5 aus y
	z := y[2:5]

	fmt.Println(w, x, y, z)
}
