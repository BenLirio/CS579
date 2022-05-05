package main

import (
  "fmt"
)



func normalize(a []int) []int {
  c := make([]int, len(a))
  for i := 0; i < len(c); i++ {
    c[i] = (a[i]+q)%q
  }
  return c
}


func mul(a []int, b []int) []int {
  c := make([]int,len(a) + len(b)-1)
  for i := 0; i < len(a); i++ {
    for j := 0; j < len(b); j++ {
      c[i+j] = (c[i+j] + a[i]*b[j])%q
    }
  }
  return normalize(c)
}

func add(a []int, b []int) []int {
  if len(a) < len(b) {
    return add(b,a)
  }
  c := make([]int, max(len(a),len(b)))
  for i := 0; i < len(c); i++ {
    c[i] = a[i]
    if i < len(b) {
      c[i] = (c[i] + b[i])%q
    }
  }
  return normalize(c)
}

func testPoly() {
  f := Poly{1,0,0,0,1}
  a := Poly{7,6,5,4,3,2,1}
  fmt.Println(a.modPoly(f))
}

func main() {
  /*
  s := []int{1,1,1,1}
  sk := make([]int, len(s))
  copy(sk,s)
  a := []int{1,1,1,1}
  e := []int{1,1,1,1}
  pk0 := mul([]int{-1},add(mul(a,s), e))
  pk1 := make([]int,len(a))
  copy(pk1,a)

  m := []int{1,1,1,1}
  u := []int{1,1,1,1}
  e1 := []int{1,1,1,1}
  e2 := []int{1,1,1,1}


  t := 2
  delta := q/t
  ct0 := add(add(mul(pk0,u),e1), mul([]int{delta},m))
  ct1 := add(mul(pk1,u),e2)
  fmt.Println(sk,ct0,ct1)
  */

  testPoly()
}
