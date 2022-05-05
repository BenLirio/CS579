package main

import (
  "./compiler"
  "./lwe"
  "os"
  "fmt"
  "log"
)

type exp = compiler.Exp

func interp(ctxt lwe.LWE, sk []int, e *exp) lwe.Cipher {
  switch e.Opt {
  case compiler.OptBinopExp:
    c1 := interp(ctxt, sk, e.Binop.E1)
    c2 := interp(ctxt, sk, e.Binop.E2)
    switch e.Binop.Binop {
    case compiler.OptAddBinop:
      return ctxt.Add(c1,c2)
    case compiler.OptMulBinop:
      return ctxt.Mul(c1,c2)
    default: fmt.Println("should not get here")
    }
  case compiler.OptBoolVal:
    return ctxt.Enc(sk,e.Bool)
  default: fmt.Println("should not get here")
  }
  log.Fatal("error")
  return lwe.Cipher{}
}

func main() {
  if len(os.Args) < 2 {
    log.Fatal("Supply equation")
  }
  e := compiler.ParseString(os.Args[1])
  ctxt := lwe.LWE{N:64,Q:1<<8,B:4}
  sk := ctxt.Gen()
  c := interp(ctxt, sk, e)
  ex := lwe.mod(c.B[lwe.M()-1] - lwe.dot(c.A[lwe.M()-1],sk))
  fmt.Println(ex)
  v := ctxt.Dec(sk, c)
  switch v {
  case true: fmt.Println("1")
  case false: fmt.Println("0")
  }
}
