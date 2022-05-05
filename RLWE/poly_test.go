package main

import (
  "testing"
)

func areSame(a Poly, b Poly) bool {
  if len(a) < len(b) { return areSame(b,a) }
  for i := 0; i < len(b); i++ {
    if a[i] != b[i] { return false }
  }
  for i := len(b); i < len(a); i++ {
    if a[i] != 0 { return false }
  }
  return true
}

func TestPolyMod(t *testing.T) {
  tdata := [][3]Poly{
    [3]Poly{
      Poly{1,0,0,0,1},
      Poly{7,6,5,4,3,2,1},
      Poly{4,4,4,4},
    },
    [3]Poly{
      Poly{1,0,0,0,1},
      Poly{1,0,0,0,1},
      Poly{0},
    },
    [3]Poly{
      Poly{1,0,0,0,1},
      Poly{1},
      Poly{1},
    },
    [3]Poly{
      Poly{1,0,1},
      Poly{1,2,1,2},
      Poly{},
    },
    [3]Poly{
      Poly{1,0,0,1},
      Poly{1,2,3,1,2,3},
      Poly{},
    },
    [3]Poly{
      Poly{1,0,1},
      Poly{1,2,1,3},
      Poly{0,-1},
    },
  }
  for _,d := range(tdata) {
    f := d[0]
    a := d[1]
    expected := d[2]
    res := a.modPoly(f)
    if len(res) < len(expected) { t.Errorf("Too few values") }
    if !areSame(expected,res) { t.Error("Expected:", expected, " Got:",res) }
  }
}
