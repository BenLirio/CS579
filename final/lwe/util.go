package lwe

import (
  "math/rand"
)

func min(a int, b int) int {
  if a < b {
    return a
  }
  return b
}

func (lwe LWE) M() int {
  return (lwe.N+1)*log2(lwe.Q)
}
func (lwe LWE) abs(a int) int {
  return min(a, lwe.Q-a)
}
func (lwe LWE) genErr() int {
  return lwe.mod(rand.Intn(2*lwe.B+1) - lwe.B)
}
func (lwe LWE) mod(a int) int {
  return ((a%lwe.Q)+lwe.Q)%lwe.Q
}

func (lwe LWE) uniformVec() []int {
  x := make([]int, lwe.N)
  for i := 0; i < lwe.N; i++ {
    x[i] = rand.Intn(lwe.Q)
  }
  return x
}

func (lwe LWE) dot(a []int, b []int) int {
  acc := 0
  for i := 0; i < lwe.N; i++ {
    acc = lwe.mod(acc + a[i]*b[i])
  }
  return acc
}


func (lwe LWE) newCipher() Cipher {
  m := (lwe.N+1)*log2(lwe.Q)
  C := Cipher{
    A: make([][]int,m),
    B: make([]int,m),
  }
  for i := 0; i < m; i++ {
    C.A[i] = make([]int,lwe.N)
  }
  return C
}
