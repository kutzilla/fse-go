package main

// Import der Format-Library für Konsolenausgaben
import "fmt"

// Initialisierung des Schachfelds
var field [8][8] bool

// Main-Methode
func main() {
  positionQueen(0)
  printField()
}

// Rekurisive Methode zur Positionierung der Damen
// n ist die Anzahl der bereits positionierten Damen
func positionQueen(n int) bool {

  // Rekursionsanfang:
  // Wenn diese Bedingung zutrifft sind 8 Damen positioniert
  // und der Alogrithmus wird beendet
  if (n == len(field)) {
    return true
  }

  // Durchlauf über alle Positionen des Feldes
  for i := 0; i < len(field); i++ {
    for j := 0; j < len(field[i]); j++ {

      // Prüft ob die aktuelle Position besetzt werden kann
      if isPositionFree(i, j) {

        // Die Dame wird positioniert
        field[j][i] = true

        // Rekusiver Aufruf der Methode für die nächste Dame
        if positionQueen(n + 1) {
          return true
        }

        // Der rekursive Aufruf war nicht erfolgreich, weshalb
        // die Dame entfernt werden muss
        field[j][i] = false
      }
    }
  }

  // Im jetzigen Zustand des Felds kann keine weitere Dame positioniert
  // werden, weshalb der laufende Rekursive Aufruf beendet wird
  return false
}

// Prüft ob sich in den Achsen zu dem übergebenen X und Y eine Dame befindet
func isPositionFree(x int, y int) bool {
  var up int
  var down int
  var left int
  var right int

  for i := 0; i < len(field); i++ {
    // Prüfung der Horizontalen und Vertikalen Achse
    if field[y][i] || field[i][x] {
      return false
    }

    // Deklaration der Bewegungsvariablen
    up = y - i
    down = y + i
    right = x + i
    left = x - i

    // Prüfung nach links Oben
    if up >= 0 && left >= 0 && field[up][left] {
      return false
    }
    // Prüfung nach rechts Oben
    if up >= 0 && right < len(field) && field[up][right] {
      return false
    }
    // Prüfung nach links Unten
    if down < len(field) && left >= 0 && field[down][left] {
      return false
    }
    // Prüfung nach rechts Unten
    if down < len(field) && right < len(field) && field[down][right] {
      return false
    }

  }
  // Falls keine Dame gefunden wurde, gilt die Postion als besetzbar
  return true
}

// Gibt das Schachfeld in der Konsole aus
func printField() {
  for i := 0; i < len(field); i++ {
    for j := 0; j < len(field[i]); j++ {
      if field[i][j] {
        fmt.Print("X  ")
      } else {
        fmt.Print("O  ")
      }
    }
    fmt.Println("")
  }
}
