package main

import (
  "fmt"
  "sort"
)

type Person struct {
  Name string
  Age int
}

type ByName []Person
func (this ByName) Len() int {return len(this) }
func (this ByName) Less(i, j int) bool { return this[i].Name < this[j].Name}
func (this ByName) Swap(i, j int) { this[i], this[j] = this[j], this[i]}

type ByAge []Person
func (this ByAge) Len() int {return len(this) }
func (this ByAge) Less(i, j int) bool { return this[i].Age < this[j].Age}
func (this ByAge) Swap(i, j int) { this[i], this[j] = this[j], this[i]}

// Methodenkopf der klassischen Schwartzian Transformation
func sortByKey(persons []Person, criteria func(p Person) bool) {
  x := criteria(persons[0])
  y := criteria(persons[1])
  fmt.Println(x)
  fmt.Println(y)
  z := x > y
}

func main() {
  var persons[3]Person
  persons[0] = Person{Name: "Willi", Age: 42}
  persons[1] = Person{Name: "Manfred", Age: 22}
  persons[2] = Person{Name: "Dieter" , Age: 63}
  fmt.Println(persons)
  sort.Sort(ByAge(persons[:]))
  fmt.Println(persons)
  sort.Sort(ByName(persons[:]))
  fmt.Println(persons)

  sortByKey(persons[:], func(p Person) interface{}{return p.Age})
}
