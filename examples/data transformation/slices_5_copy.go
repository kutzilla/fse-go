package main

import "fmt"

func main() {

	x := []int{1, 2, 3}
	y := make([]int, 2)

	// kopiere slice x nach slice y
	copy(y, x)
	fmt.Println(x, y)

}
