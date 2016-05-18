package main

import (
	"fmt"
  "sort"
  "time"
	"github.com/danverbraganza/keysort"
)

var count int

// Struktur von Personen
type Person struct {
	Name string
	Age  int
}

// Simulation einer teuren Operation
func (this *Person) getAge() int {
  count++
  time.Sleep(100 * time.Millisecond)
  return this.Age
}

// Struktur von Personen Array
type ByAge []Person

// Methoden die von dem interface sort.Interface zur Sortierung benötigt werden
func (this ByAge) Len() int           { return len(this) }
func (this ByAge) Less(i, j int) bool { return this[i].getAge() < this[j].getAge() }
func (this ByAge) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
// Methoden die speziell von keysort.Interface gebraucht werden
func (this ByAge) Key(i int) (interface{}, error) { return this[i].getAge(), nil }
func (this ByAge) LessVal(i, j interface{}) bool {return i.(int) < j.(int)}


// Hilfsmethode um die Geschwindigkeit der beiden Sortierungen zu vergleichen
func BenchmarkSortFunc(factory func() sort.Interface) {
	count = 0
	start := time.Now()
	s := factory()
	sort.Sort(s)
	end := time.Now()
	fmt.Println("sorted:", s)
	fmt.Println("Calls to Keyfunc:", count)
	fmt.Println("Elapsed time:", end.Sub(start))
}

func main() {
  // Slice mit Personen
	persons := []Person{
		{Name: "Willi", Age: 42},
		{Name: "Manfred", Age: 61},
		{Name: "Dieter", Age: 63},
    {Name: "Dave", Age: 27},
    {Name: "Matthias", Age: 22},
    {Name: "Marcel", Age: 27},
	}

  // Aufruf der beiden Sortierfunktionen
  BenchmarkSortFunc(func() sort.Interface { return ByAge(persons) })
  BenchmarkSortFunc(func() sort.Interface { return keysort.Keysort(ByAge(persons)) })
}
