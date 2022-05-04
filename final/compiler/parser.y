%{

package compiler

import (
)

%}

%union {
	exp *Exp
}

%type	<exp>	expr expr1 expr2

%token PLUS
%token DASH
%token STAR
%token FSLASH
%token LPAREN
%token RPAREN
%token <exp> ZERO
%token <exp> ONE

%%

top: expr { expOut = $1 }


expr: expr1
     | expr PLUS expr1 { $$ = NewBinopExp(OptAddBinop,$1,$3) }

expr1: expr2
     | expr1 STAR expr2 { $$ = NewBinopExp(OptMulBinop,$1,$3) }

expr2: ZERO
     | ONE
     | LPAREN expr RPAREN { $$ = $2 }


%%
