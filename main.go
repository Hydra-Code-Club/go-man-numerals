package main

import (
  "fmt"
  s "strings"
  "sort"
)

type numeral struct {
  value int
}

func getCharToValueTable() map[byte]int {
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
func getValueToCharTable() map[int]string {
  m := getCharToValueTable()
  n := make(map[int]string)
  for k, v := range m {
      n[v] = string(k)
  }
  n[4] = "IV"
  n[9] = "IX"
  n[40] = "XL"
  n[400] = "CD"
  n[900] = "CM"
  return n
}

// TODO: map keys helper function (templates!)

// TODOn't: validating helper function

func (n numeral) getNumber() int {
  return n.value
}

func (n numeral) toString() string {
  charTable := getValueToCharTable()
  // get keys
  keys := make([]int, len(charTable))
  i := 0
  for k := range charTable {
    keys[i] = k
    i++
  }
  // sort
  sort.Ints(keys)
  // iterate
  num := n.getNumber()
  str := ""
  for i := len(keys) - 1; i >= 0; i-- { // RTL
    for num >= keys[i] {
      num -= keys[i]
      str += string(charTable[keys[i]])
    }
  }
  return str
}

func parseString(str string) numeral {
  digitTable := getCharToValueTable()
  upper := s.ToUpper(str)
  sum := 0
  highest := 0
  for i := len(upper) - 1; i >= 0; i-- { // RTL
    val := digitTable[upper[i]]
    if val >= highest {
      sum += val
      highest = val
    } else {
      sum -= val
    }
  }

  // TIL: You don't need to specify key. Set by order.
  return numeral{value: sum}
}

func main() {
  fmt.Println("Hello, Anna!")

  num := parseString("XIV")

  fmt.Println(num.getNumber())

  fmt.Println(num.toString())
}
