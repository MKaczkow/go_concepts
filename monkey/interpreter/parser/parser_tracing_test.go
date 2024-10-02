package parser

import (
	"testing"
)

func TestIdentLevel(t *testing.T) {
	// Test ident level at various levels
	tests := []struct {
		level       int
		expectedOut string
	}{
		{1, ""},       // No indentation at level 1
		{2, "\t"},     // One tab at level 2
		{3, "\t\t"},   // Two tabs at level 3
		{4, "\t\t\t"}, // Three tabs at level 4
	}

	for _, test := range tests {
		traceLevel = test.level
		out := identLevel()
		if out != test.expectedOut {
			t.Errorf("For traceLevel %d, expected %q, got %q", test.level, test.expectedOut, out)
		}
	}
}

func TestIncDecIdent(t *testing.T) {
	// Reset traceLevel
	traceLevel = 0

	// Increment and check
	incIdent()
	if traceLevel != 1 {
		t.Errorf("Expected traceLevel 1, got %d", traceLevel)
	}

	// Increment again and check
	incIdent()
	if traceLevel != 2 {
		t.Errorf("Expected traceLevel 2, got %d", traceLevel)
	}

	// Decrement and check
	decIdent()
	if traceLevel != 1 {
		t.Errorf("Expected traceLevel 1 after decrement, got %d", traceLevel)
	}

	// Decrement again and check
	decIdent()
	if traceLevel != 0 {
		t.Errorf("Expected traceLevel 0 after second decrement, got %d", traceLevel)
	}
}
