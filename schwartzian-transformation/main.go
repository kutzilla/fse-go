package main

import (
  "fmt"
  "sort"
)

type Person struct {
  Name string
  Age int
}

//Typ zur Sortierung nach Namen, implementiert sort.Interface
type ByName []Person
func (this ByName) Len() int {return len(this) }
func (this ByName) Less(i, j int) bool { return this[i].Name < this[j].Name}
func (this ByName) Swap(i, j int) { this[i], this[j] = this[j], this[i]}

//Typ zur Sortierung nach Alter, implementiert sort.Interface
type ByAge []Person
func (this ByAge) Len() int {return len(this) }
func (this ByAge) Less(i, j int) bool { return this[i].Age < this[j].Age}
func (this ByAge) Swap(i, j int) { this[i], this[j] = this[j], this[i]}

func main() {
  //Array mit drei Personen
  var persons[3]Person
  persons[0] = Person{Name: "Willi", Age: 42}
  persons[1] = Person{Name: "Manfred", Age: 22}
  persons[2] = Person{Name: "Dieter" , Age: 63}
  fmt.Println(persons)
  //Aufruf der Sortierfunktion von der sort Bibliothek
  //Übergabe von personen als Slice
  sort.Sort(ByAge(persons[:]))
  fmt.Println(persons)
  
  sort.Sort(ByName(persons[:]))
  fmt.Println(persons)

}
