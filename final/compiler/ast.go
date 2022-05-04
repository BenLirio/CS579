package compiler

import (
  "fmt"
)

// Exp
const (
  OptBinopExp int = 1
  OptBoolVal int = 2
)
type Exp struct {
  Opt int
  Binop *BinopExp
  Bool bool
}
func (e *Exp) String() string {
  switch e.Opt {
  case OptBoolVal: if e.Bool { return "1" } else { return "0" }
  case OptBinopExp:
    return fmt.Sprintf("%s + %s", e.Binop.E1.String(), e.Binop.E2.String())
  default: return "Not Implemented"
  }
}

// Binop
type BinopExp struct {
  Binop int
  E1 *Exp
  E2 *Exp
}
const (
  OptAddBinop int = 1
  OptMulBinop int = 2
)
func NewBinopExp(bop int, e1 *Exp, e2 *Exp) *Exp {
  return &Exp{
    Opt: OptBinopExp,
    Binop: &BinopExp{
      Binop: bop,
      E1: e1,
      E2: e2,
    },
  }
}

// Bool
func NewBoolVal(a bool) *Exp {
  return &Exp{
    Opt: OptBoolVal,
    Bool: a,
  }
}
