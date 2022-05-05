package main

import (
  "math/rand"
)

func (rlwe RLWE) mul(a []int, b []int) []int {
  c := make([]int, rlwe.N)
  for i := 0; i < rlwe.N; i++ {
    for j := 0; j < rlwe.N; j++ {
      idx := i + j
      if idx < rlwe.N {
        c[idx] += a[i]*b[j]
      } else {
        idx -= rlwe.N
        c[idx] -= a[i]*b[j]
      }
    }
  }
  return c
}
func (rlwe RLWE) add(a []int, b []int) []int {
  c := make([]int, rlwe.N)
  for i := 0; i < rlwe.N; i++ {
    c[i] = a[i]+b[i]
  }
  return c
}
func (rlwe RLWE) neg(a []int) []int {
  c := make([]int, rlwe.N)
  for i := 0; i < rlwe.N; i++ {
    c[i] = -a[i]
  }
  return c
}

func (rlwe RLWE) chi() int {
  // Add security param to distribution
  x := rand.NormFloat64()
  sign := 1
  if x < 0 {
    sign = -1
    x = -x
  }
  xi := int(x)
  return xi*sign
}

func (rlwe RLWE) unifq() int {
  return rand.Intn(rlwe.Q)
}
func (rlwe RLWE) unifqn() []int {
  x := make([]int, rlwe.N)
  for i := 0; i < rlwe.N; i++ {
    x[i] = rlwe.unifq()
  }
  return x
}

func (rlwe RLWE) chin() []int {
  x := make([]int, rlwe.N)
  for i := 0; i < rlwe.N; i++ {
    x[i] = rlwe.chi()
  }
  return x
}
