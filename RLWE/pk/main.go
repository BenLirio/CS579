package main

import (
  "fmt"
  "time"
  "math/rand"
)

type RLWE struct {
  N int
  Q int
  T int
  delta int
}

func newRLWE(n int, q int, t int) RLWE {
  return RLWE{
    N: n,
    Q: q,
    T: t,
    delta: q/t,
  }
}



func main() {
  rand.Seed(time.Now().UnixNano())
  rlwe := newRLWE(1<<3,32,16)
  s := rlwe.chin()
  a := rlwe.unifqn()
  c0 := a
  e := rlwe.chin()
  c1 := rlwe.add(rlwe.mul(s,a), e)
  xe := rlwe.add(c1, rlwe.neg(rlwe.mul(s,c0)))
  fmt.Println(xe)
}
