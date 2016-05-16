package main

import "fmt"

func main() {
	var a []int

	// append bei einem Null-Slice
	x := append(a, 9)

	// das Slice wächst eigenständig
	y := append(x, 1)

	// auch Kommaseperiert
	z := append(y, 2, 3, 4)

	fmt.Println(x, y, z)
}
