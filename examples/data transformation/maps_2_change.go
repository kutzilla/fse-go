package main
import "fmt"

func main() {
	m := make(map[int]string)

	// Wert zuweisen
	m[11] = "Reus"
  m[17] = "Aubameyang"
	m[15] = "Hummels"

  // Wert auslesen
	reus := m[11]
	fmt.Println("Spieler mit der Nr. 11: ", reus)

	// Wert löschen
	delete(m, 15)

	v15, ok15 := m[15]
	fmt.Println("Spieler mit der Nr. 15 vorhanden: ",v15, ok15)

	v17, ok17 := m[17]
	fmt.Println("Spieler mit der Nr. 17 vorhanden: ",v17, ok17)
}
