package main

import (
  "math/rand"
  "./lwe"
  "fmt"
)

func f(ctxt lwe.LWE, sk []int, n int) (lwe.Cipher,bool) {
  if n == 0 {
    v := false
    if rand.Intn(2) == 0 {
      v = true
    }
    return ctxt.Enc(sk,v),v
  }
  lc,lt := f(ctxt,sk,n-1)
  rc,rt := f(ctxt,sk,n-1)
  return ctxt.Mul(lc,rc),lt&&rt
}

func main() {
  ctxt := lwe.LWE{N:64,Q:1<<8,B:4}
  sk := ctxt.Gen()
  depth := 0
  for {
    c,t := f(ctxt,sk,depth)
    fmt.Printf("[ciphers=%d]\n", 1<<depth)
    if ctxt.Dec(sk,c) != t {
      break
    }
    depth += 1
  }
}
