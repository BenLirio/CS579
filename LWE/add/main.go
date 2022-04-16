package main

import (
  "math/rand"
  "fmt"
  "time"
)

type LWE struct {
  N int
  Q int
  B int
}
type Message bool
type Vec []int
type Cipher struct {
  c0 []int
  c1 int
  eb int
}

func min(a int, b int) int {
  if a < b {
    return a
  }
  return b
}

func (lwe LWE) mod(a int) int {
  return ((a%lwe.Q)+lwe.Q)%lwe.Q
}
func (lwe LWE) uniformVec() Vec {
  x := make(Vec, lwe.N)
  for i := 0; i < lwe.N; i++ {
    x[i] = rand.Intn(lwe.Q)
  }
  return x
}
func (lwe LWE) genErr() int {
  return lwe.mod(rand.Intn(2*lwe.B+1) - lwe.B)
}
func (lwe LWE) abs(a int) int {
  return min(a, lwe.Q-a)
}
func (lwe LWE) add(a Vec, b Vec) Vec {
  c := make(Vec, lwe.N)
  for i := 0; i < lwe.N; i++ {
    c[i] = lwe.mod(a[i] + b[i])
  }
  return c
}
func (lwe LWE) dot(a Vec, b Vec) int {
  acc := 0
  for i := 0; i < lwe.N; i++ {
    acc = lwe.mod(acc + a[i]*b[i])
  }
  return acc
}


func (lwe LWE) Gen() Vec {
  return lwe.uniformVec()
}

func (lwe LWE) Enc(s Vec, x Message) Cipher {
  a := lwe.uniformVec()
  e := lwe.genErr()
  c1 := lwe.mod(lwe.dot(a,s) + e)
  if x {
    c1 = lwe.mod(c1 + lwe.Q/2)
  }
  return Cipher{
    c0: a,
    c1: c1,
    eb: lwe.B + 1,
  }
}

func (lwe LWE) Dec(s Vec, c Cipher) Message {
  ex := lwe.mod(c.c1 - lwe.dot(c.c0, s))
  return lwe.abs(ex) > lwe.Q/4
}

func (lwe LWE) Add(a Cipher, b Cipher) Cipher {
  return Cipher{
    c0: lwe.add(a.c0,b.c0),
    c1: lwe.mod(a.c1+b.c1),
    eb: a.eb + b.eb + 1,
  }
}

func randMessage() Message {
  if rand.Intn(2) == 1 {
    return Message(true)
  }
  return Message(false)
}

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  lwe := LWE{
    N: 128,
    Q: 64,
    B: 2,
  }
  trials := 1000
  correct := 0
  for i := 0; i < trials; i++ {
    s := lwe.Gen()
    x0 := randMessage()
    x1 := randMessage()
    c0 := lwe.Enc(s, x0)
    c1 := lwe.Enc(s, x1)

    c2 := lwe.Add(c0,c1)
    x2 := lwe.Dec(s,c2)
    ok := false
    if (x0 && x1) || (!x0 && !x1) {
      if x2 == false {
        ok = true
      }
    } else {
      if x2 == true {
        ok = true
      }
    }
    if ok {
      correct += 1
    }
  }
  fmt.Printf("N: %d, Q: %d, B: %d\n", lwe.N, lwe.Q, lwe.B)
  fmt.Printf("%d/%d\n",correct,trials)
}
