package compiler
import (
	"log"
	"unicode/utf8"
)

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const eof = 0


// The parser uses the type <prefix>Lex as a lexer. It must provide
// the methods Lex(*<prefix>SymType) int and Error(string).
type exprLex struct {
	line []byte
	peek rune
}

// The parser calls this method to get each new token. This
// implementation returns operators and NUM.
func (x *exprLex) Lex(yylval *exprSymType) int {
	for {
		c := x.next()
		switch c {
		case eof: return eof
    case '0': return x.zero(yylval)
    case '1': return x.one(yylval)
    case '+': return PLUS
    case '*': return STAR
    case '(': return LPAREN
    case ')': return RPAREN
		case ' ', '\t', '\n', '\r': continue
		default: log.Printf("unrecognized character %q", c)
		}
	}
}
func (x *exprLex) zero(yylval *exprSymType) int {
  yylval.exp = NewBoolVal(false)
	return ZERO
}
func (x *exprLex) one(yylval *exprSymType) int {
  yylval.exp = NewBoolVal(true)
	return ONE
}

// Return the next rune for the lexer.
func (x *exprLex) next() rune {
	if x.peek != eof {
		r := x.peek
		x.peek = eof
		return r
	}
	if len(x.line) == 0 {
		return eof
	}
	c, size := utf8.DecodeRune(x.line)
	x.line = x.line[size:]
	if c == utf8.RuneError && size == 1 {
		log.Print("invalid utf8")
		return x.next()
	}
	return c
}

// The parser calls this method on a parse error.
func (x *exprLex) Error(s string) {
	log.Printf("parse error: %s", s)
}
