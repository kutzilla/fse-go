package main
import "fmt"
func main() {
	# String Array
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	#1. Float Array
	x := [5]float64{ 1, 2, 3, 4, 5 }
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	#2. Float Array
	y := [3]float64{
		97,
		98,
		99,
	}
	fmt.Println(y)
}
