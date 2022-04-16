package main

import (
  "../lwe"
  "fmt"
  "math/rand"
)

func main() {
  ctxt := lwe.LWE{N: 8, Q: 128, B: 1}
  s := ctxt.Gen()
  //zero := ctxt.Enc(s,false)
  one := ctxt.Enc(s,true)
  trials := 100
  correct := 0
  count := 0
  for count < trials {
    count += 1
    r := rand.Intn(1<<3)
    v0 := r&1 == 1
    cv0 := ctxt.Enc(s, v0)
    v1 := r&2 == 2
    cv1 := ctxt.Enc(s, v1)
    idx := r&4 == 4
    cidx := ctxt.Enc(s,idx)

    cout := ctxt.Add(ctxt.Mul(cidx,cv1),ctxt.Mul(ctxt.Add(cidx,one), cv0))
    out := ctxt.Dec(s,cout)

    expected := (idx&&v1)||((!idx)&&v0)
    if expected == out {
      correct += 1
    } else {
    }
  }
  fmt.Printf("%d/%d\n", correct, trials)
}
