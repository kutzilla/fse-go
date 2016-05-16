package main
import "fmt"
//variadische funktion,benannter Rückgabewert sum
func summiere(args ...int)  (sum int) {
  for _, value:= range args{
    sum+=value
  }
  return
}
//subtrahieren funktion rückgabe int
func subtrahiere(minuend,subtrahend int) int{
  return minuend - subtrahend
}
//zwei Rückgabewerte
func dividiere(divident, divisor float64) (float64,bool){
  if(divisor != 0){
    return divident/divisor, true
  }
    return 0,false
}
func main() {
  fmt.Println(summiere(1,2,3,4))
  fmt.Println(subtrahiere(1,2))
  quot,ok := dividiere(2,0)
  fmt.Println(quot,ok)
}
