package token

import (
	"testing"
)

func TestLookupIdent(t *testing.T) {
	tests := []struct {
		input    string
		expected TokenType
	}{
		// Test for keywords
		{"fn", FUNCTION},
		{"let", LET},
		{"true", TRUE},
		{"false", FALSE},
		{"if", IF},
		{"else", ELSE},
		{"return", RETURN},

		// Test for identifiers (not keywords)
		{"myVariable", IDENT},
		{"x", IDENT},
		{"add", IDENT},
	}

	for _, tt := range tests {
		tokType := LookupIdent(tt.input)
		if tokType != tt.expected {
			t.Errorf("LookupIdent(%q) = %q; want %q", tt.input, tokType, tt.expected)
		}
	}
}

func TestTokenStruct(t *testing.T) {
	// Test creating a token for an identifier
	tok := Token{Type: IDENT, Literal: "myVariable"}
	if tok.Type != IDENT {
		t.Errorf("Token.Type = %q; want %q", tok.Type, IDENT)
	}
	if tok.Literal != "myVariable" {
		t.Errorf("Token.Literal = %q; want %q", tok.Literal, "myVariable")
	}

	// Test creating a token for an integer literal
	tok = Token{Type: INT, Literal: "12345"}
	if tok.Type != INT {
		t.Errorf("Token.Type = %q; want %q", tok.Type, INT)
	}
	if tok.Literal != "12345" {
		t.Errorf("Token.Literal = %q; want %q", tok.Literal, "12345")
	}
}
