package test_test

import (
	"lpp/lpp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetStatement(t *testing.T) {
	program := lpp.NewProgram([]lpp.Stmt{
		lpp.LetStatement{
			Token: lpp.Token{
				Token_type: lpp.LET,
				Literal:    "var",
			},
			Name: lpp.NewIdentifier(lpp.Token{
				Token_type: lpp.IDENT,
				Literal:    "mi_var",
			}, "mi_var"),
			Value: lpp.NewIdentifier(lpp.Token{
				Token_type: lpp.IDENT,
				Literal:    "otra_var",
			}, "otra_var"),
		},
	})

	assert.Equal(t, "var mi_var = otra_var;", program.Str())
}

func TestReturnStatements(t *testing.T) {
	program := lpp.NewProgram([]lpp.Stmt{
		lpp.LetStatement{
			Token: lpp.Token{
				Token_type: lpp.LET,
				Literal:    "var",
			},
			Name: lpp.NewIdentifier(lpp.Token{
				Token_type: lpp.IDENT,
				Literal:    "x",
			}, "x"),
			Value: lpp.NewIdentifier(lpp.Token{
				Token_type: lpp.INT,
				Literal:    "5",
			}, "5"),
		},
		lpp.ReturnStament{
			Token: lpp.Token{
				Token_type: lpp.RETURN,
				Literal:    "regresa",
			},
			ReturnValue: lpp.NewIdentifier(
				lpp.Token{
					Token_type: lpp.IDENT,
					Literal:    "x",
				}, "x"),
		},
	})

	assert.Equal(t, "var x = 5; regresa x;", program.Str())
}

func TestIntegerExpression(t *testing.T) {
	value := 5
	program := lpp.NewProgram([]lpp.Stmt{
		lpp.LetStatement{
			Token: lpp.Token{
				Token_type: lpp.LET,
				Literal:    "var",
			},
			Name: lpp.NewIdentifier(lpp.Token{
				Token_type: lpp.IDENT,
				Literal:    "x",
			}, "x"),
			Value: lpp.NewInteger(lpp.Token{
				Token_type: lpp.INT,
				Literal:    "5",
			}, &value),
		},
	})

	assert.Equal(t, "var x = 5;", program.Str())
}