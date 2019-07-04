package ast

import (
	"goRamble/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar;" {
		t.Errorf("program.String( wrong. got=%q", program.String())
	}
}
