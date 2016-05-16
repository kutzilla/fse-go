package main

import (
	"fmt"
)

func main() {
	fmt.Println("Fahrt los...")
	for i := 1; i <= 3; i++ {
		// if-else Variante
		if i == 1 {
			fmt.Println("eins da")
		} else if i == 2 {
			fmt.Println("zwei da ")
		} else if i == 3 {
			defer fmt.Println("drei hat Verspätung")
		}
	}
	k := 4
	//while Schleife
	for k <= 7 {
		//switch case Variante
		switch k {
		case 4:
			fmt.Println("vier da")
		case 5:
			fmt.Println("fünf da")
		case 6:
			fmt.Println("sechs da")
		default:
			fmt.Println("wer ist das denn?")
		}
		k++
	}
}
