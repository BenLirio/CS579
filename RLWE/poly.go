package main

var q int = 5
var f []int = []int{
  1, // x^0
  0, // x^1
  0, // x^2
  0, // x^3
  1, // x^4
}
type Poly []int

func mod(a int, m int) int {
  return ((a%m)+m)%m
}

func (a Poly) _cmp(b Poly) int {
  // Not normalized
  n := max(len(a),len(b))
  for i := n-1; i >= 0; i-- {
    aVal := 0
    bVal := 0
    if i < len(a) {
      aVal = a[i]
    }
    if i < len(b) {
      bVal = b[i]
    }
    if aVal > bVal { return 1 }
    if bVal > aVal { return -1 }
  }
  return 0
}

func (a Poly) _gt(b Poly) bool { return a._cmp(b) == 1 }
func (a Poly) _eq(b Poly) bool { return a._cmp(b) == 0 }
func (a Poly) _lt(b Poly) bool { return !a._eq(b) && !a._gt(b) }
func (a Poly) _lte(b Poly) bool { return !a._gt(b) }
func (a Poly) _gte(b Poly) bool { return !a._lt(b) }

func (a Poly) _neg() Poly {
  c := make(Poly, len(a))
  copy(c,a)
  for i := 0; i < len(a); i++ {
    c[i] = -c[i]
  }
  return c
}

func (a Poly) _add(b Poly) Poly {
  n := max(len(a),len(b))
  c := make(Poly, n)
  for i := 0; i < n; i++ {
    aVal := 0
    bVal := 0
    if i < len(a) { aVal = a[i] }
    if i < len(b) { bVal = b[i] }
    c[i] = aVal + bVal
  }
  return c
}

func (a Poly) _mulx() Poly {
  c := make(Poly,len(a)+1)
  for i := 1; i < len(c); i++ {
    c[i] = a[i-1]
  }
  return c
}
func (a Poly) _divx() Poly {
  if a[0] != 0 { panic("Div will be wrong") }
  c := make(Poly,len(a)-1)
  for i := 0; i < len(c); i++ {
    c[i] = a[i+1]
  }
  return c
}

func (a Poly) modPoly(f Poly) Poly {
  c := make(Poly,len(a))
  copy(c,a)
  for c._gte(f) {
    cur := make(Poly,len(f))
    copy(cur,f)
    for cur._lte(c) {
      cur = cur._mulx()
    }
    cur = cur._divx()
    for c._gte(cur) {
      c = c._add(cur._neg())
    }
  }
  return c
}

func (a Poly) modCoeff(m int) Poly {
  c := make([]int, len(a))
  copy(c,a)
  for i := 0; i < len(c); i++ {
    c[i] = mod(c[i],m)
  }
  return c
}

func (a Poly) mul(b Poly) Poly {
  return b
}

func max(a int, b int) int {
  if a > b {
    return a
  }
  return b
}
func min(a int, b int) int {
  if a < b {
    return a
  }
  return b
}
