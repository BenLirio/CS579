package main

import (
  "fmt"
  "math/rand"
)

// Always hold the invariants
// Allow some casts
// Functions starting with _ do not hold invariants
type LWE struct {
  N int
  Q int
  B int
}

type Cipher struct {
  A [][]int
  B []int
}

func log2(a int) int {
  if a == 0 { return 0 }
  ct := 0
  for a != 1 {
    ct += 1
    a = a>>1
  }
  return ct
}

func (lwe LWE) Gen() []int {
  return lwe.uniformVec()
}

func (lwe LWE) G(C Cipher) {
  // Does not modify all values
  // So inplace is more efficient
  for i := 0; i < lwe.N; i++ {
    for j := 0; j < log2(lwe.Q); j++ {
      col := i*log2(lwe.Q)+j
      C.A[col][i] = lwe.mod(C.A[col][i]+(1<<j))
    }
  }
  for j := 0; j < log2(lwe.Q); j++ {
    idx := lwe.N*log2(lwe.Q)+j
    C.B[idx] = lwe.mod(C.B[idx]+(1<<j))
  }
}

func (lwe LWE) Enc(s []int, x bool) Cipher {
  C := lwe.newCipher()
  for i := 0; i < lwe.M(); i++ {
    C.A[i] = lwe.uniformVec()
    e := lwe.genErr()
    C.B[i] = lwe.mod(lwe.dot(C.A[i],s) + e)
  }
  if x { lwe.G(C) }
  return C
}

func (lwe LWE) Dec(s []int, C Cipher) bool {
  ex := lwe.mod(C.B[lwe.M()-1] - lwe.dot(C.A[lwe.M()-1],s))
  return lwe.abs(ex) > lwe.Q/4
}

func (lwe LWE) Add(C0 Cipher, C1 Cipher) Cipher {
  C := lwe.newCipher()
  for i := 0; i < lwe.M(); i++ {
    for j := 0; j < lwe.N; j++ {
      C.A[i][j] = lwe.mod(C0.A[i][j]+C1.A[i][j])
    }
  }
  for i := 0; i < lwe.M(); i++ {
    C.B[i] = lwe.mod(C0.B[i]+C1.B[i])
  }
  return C
}

func (lwe LWE) Mul(C0 Cipher, C1 Cipher) Cipher {
  R := make([][]int,lwe.M())

  for i := 0; i < lwe.M(); i++ {
    R[i] = make([]int,lwe.M())
    for j := 0; j < lwe.M(); j++ {
      if i/log2(lwe.Q) < lwe.N {
        R[i][j] = (C1.A[j][i/log2(lwe.Q)]&(1<<(i%log2(lwe.Q))))>>(i%log2(lwe.Q))
      } else {
        R[i][j] = (C1.B[j]&(1<<(i%log2(lwe.Q))))>>(i%log2(lwe.Q))
      }
    }
  }
  L := make([][]int,lwe.N+1)
  for i := 0; i < lwe.N+1; i++ {
    L[i] = make([]int,lwe.M())
    for j := 0; j < lwe.M(); j++ {
      if i < lwe.N {
        L[i][j] = C0.A[j][i]
      } else {
        L[i][j] = C0.B[j]
      }
    }
  }
  C := lwe.newCipher()
  for i := 0; i < lwe.N+1; i++ {
    for j := 0; j < lwe.M(); j++ {
      acc := 0
      for k := 0; k < lwe.M(); k++ {
        acc += L[i][k]*R[k][j]
      }
      if i < lwe.N {
        C.A[j][i] = lwe.mod(acc)
      } else {
        C.B[j] = lwe.mod(acc)
      }
    }
  }
  return C
}

func main() {
  lwe := LWE{
    N: 8,
    Q: 64,
    B: 1,
  }
  trials := 100
  correct := 0
  for i := 0; i < trials; i++ {
    s := lwe.Gen()
    r := rand.Intn(4)
    x0 := (r&1) == 1
    x1 := (r&2) == 2
    C0 := lwe.Enc(s,x0)
    C1 := lwe.Enc(s,x1)
    C2 := lwe.Enc(s,x1)
    if lwe.Dec(s,lwe.Mul(lwe.Add(C0,C1),C2)) == ((lwe.Dec(s,C0) != lwe.Dec(s,C1)) && lwe.Dec(s,C2)) {
      correct += 1
    }
  }
  fmt.Printf("N: %d, Q: %d, B: %d\n", lwe.N, lwe.Q, lwe.B)
  fmt.Printf("%d/%d\n",correct,trials)
}
