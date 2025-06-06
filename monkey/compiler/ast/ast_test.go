package ast

import (
	"monkey/compiler/token"
	"testing"
)

func TestString(t *testing.T) {

	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestExpressionStatementString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&ExpressionStatement{
				Token: token.Token{Type: token.IDENT, Literal: "foo"},
				Expression: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "foo"},
					Value: "foo",
				},
			},
		},
	}

	expected := "foo"
	if program.String() != expected {
		t.Errorf("program.String() wrong. got=%q, want=%q", program.String(), expected)
	}
}

func TestIntegerLiteralString(t *testing.T) {
	il := &IntegerLiteral{
		Token: token.Token{Type: token.INT, Literal: "5"},
		Value: 5,
	}

	expected := "5"
	if il.String() != expected {
		t.Errorf("IntegerLiteral.String() wrong. got=%q, want=%q", il.String(), expected)
	}
}

func TestPrefixExpressionString(t *testing.T) {
	pe := &PrefixExpression{
		Token:    token.Token{Type: token.BANG, Literal: "!"},
		Operator: "!",
		Right: &IntegerLiteral{
			Token: token.Token{Type: token.INT, Literal: "5"},
			Value: 5,
		},
	}

	expected := "(!5)"
	if pe.String() != expected {
		t.Errorf("PrefixExpression.String() wrong. got=%q, want=%q", pe.String(), expected)
	}
}

func TestInfixExpressionString(t *testing.T) {
	ie := &InfixExpression{
		Token:    token.Token{Type: token.PLUS, Literal: "+"},
		Left:     &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
		Operator: "+",
		Right:    &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "10"}, Value: 10},
	}

	expected := "(5 + 10)"
	if ie.String() != expected {
		t.Errorf("InfixExpression.String() wrong. got=%q, want=%q", ie.String(), expected)
	}
}

func TestCallExpressionString(t *testing.T) {
	ce := &CallExpression{
		Token: token.Token{Type: token.LPAREN, Literal: "("},
		Function: &Identifier{
			Token: token.Token{Type: token.IDENT, Literal: "add"},
			Value: "add",
		},
		Arguments: []Expression{
			&Identifier{
				Token: token.Token{Type: token.IDENT, Literal: "x"},
				Value: "x",
			},
			&Identifier{
				Token: token.Token{Type: token.IDENT, Literal: "y"},
				Value: "y",
			},
		},
	}

	expected := "add(x, y)"
	if ce.String() != expected {
		t.Errorf("CallExpression.String() wrong. got=%q, want=%q", ce.String(), expected)
	}
}
