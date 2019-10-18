package main

import (
  "fmt"
  s "strings"
)

type numeral struct {
  value int
}

func getDigitMap() map[byte]int {
  // TODO: instance
  digits := make(map[byte]int)
  digits['I'] = 1
  digits['V'] = 5
  digits['X'] = 10
  digits['L'] = 50
  digits['C'] = 100
  digits['D'] = 500
  digits['M'] = 1000
  return digits
}

func (n numeral) getNumber() int {
  return n.value
}
/*
func (n *numeral) toString() string {
  return n.value
}
*/
func parseString(str string) numeral {
  digits := getDigitMap()
  upper := s.ToUpper(str)
  sum := 0
  highest := 0
  for i := len(upper) - 1; i >= 0; i-- { // RTL
    val := digits[upper[i]]
    if val >= highest {
      sum += val
      highest = val
    } else {
      sum -= val
    }
  }

  return numeral{value: sum}
}

func main() {
  fmt.Println("Hello, Anna!")

  num := parseString("XIV")

  fmt.Println(num.getNumber())
}
