package main

import "testing"

func parseTest(t *testing.T, str string, expected int) {
  tmp := parseString(str).getNumber()
  if (tmp != expected) {
    t.Errorf("%s not %d, was %d", str, expected, tmp)
  }
}

// TIL: Capitalized vars and functions are exported from the package
func TestParsing(t *testing.T) {
  parseTest(t, "X", 10)
  parseTest(t, "XI", 11)
  parseTest(t, "IX", 9)
  parseTest(t, "MCMLXXVII", 1977)
  parseTest(t, "IIIIIIIIX", 2)
}

func encodeTest(t *testing.T, num int, expected string) {
  tmp := numeral{num}.toString()
  if (tmp != expected) {
    t.Errorf("%d not %s, was %s", num, expected, tmp)
  }
}
func TestEncoding(t *testing.T) {
  encodeTest(t, 10, "X")
  encodeTest(t, 11, "XI")
  encodeTest(t, 9, "IX")
  encodeTest(t, 1977, "MCMLXXVII")
  encodeTest(t, 2, "II")
}
