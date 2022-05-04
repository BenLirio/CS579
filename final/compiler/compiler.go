// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file holds the go generate command to run yacc on the grammar in expr.y.
// To build expr:
//	% go generate
//	% go build

//go:generate ragel -o lex.go -Z lex.rl 
//go:generate goyacc -o parser.go -p expr parser.y

// Expr is a simple expression evaluator that serves as a working example of
// how to use Go's yacc implementation.
package compiler

import (
	"bufio"
	"io"
	"log"
	"os"
)

var expOut *Exp
func ParseString() *Exp {
	in := bufio.NewReader(os.Stdin)
  if _, err := os.Stdout.WriteString("> "); err != nil {
    log.Fatalf("WriteString: %s", err)
  }
  line, err := in.ReadBytes('\n')
  if err == io.EOF {
    return expOut
  }
  if err != nil {
    log.Fatalf("ReadBytes: %s", err)
  }

  exprParse(&exprLex{line: line})
  return expOut
}
