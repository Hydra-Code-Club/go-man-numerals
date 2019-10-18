package main

import "testing"

func parseTest(t *testing.T, str string, expected int) {
  tmp := parseString(str).getNumber()
  if (tmp != expected) {
    t.Errorf("%s not %d, was %d", str, expected, tmp)
  }
}

func TestTesting(t *testing.T) {
  parseTest(t, "X", 10)
  parseTest(t, "XI", 11)
  parseTest(t, "IX", 9)
  parseTest(t, "MCMLXXVII", 1977)
  parseTest(t, "IIIIIIIIX", 2)
}
