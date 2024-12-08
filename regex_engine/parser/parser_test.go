package parser

import (
	"reflect"
	"testing"
)

// func TestParseLiteral(t *testing.T) {
// 	regex := "a"
// 	ctx := parse(regex)

// 	expectedTokens := []token{
// 		{tokenType: literal, value: 'a'},
// 	}

// 	if !reflect.DeepEqual(ctx.tokens, expectedTokens) {
// 		t.Errorf("Expected tokens %v, got %v", expectedTokens, ctx.tokens)
// 	}
// }

// func TestParseGroup(t *testing.T) {
// 	regex := "(ab)"
// 	ctx := parse(regex)

// 	expectedTokens := []token{
// 		{
// 			tokenType: group,
// 			value: []token{
// 				{tokenType: literal, value: 'a'},
// 				{tokenType: literal, value: 'b'},
// 			},
// 		},
// 	}

// 	if !reflect.DeepEqual(ctx.tokens, expectedTokens) {
// 		t.Errorf("Expected tokens %v, got %v", expectedTokens, ctx.tokens)
// 	}
// }

func TestParseBracket(t *testing.T) {
	regex := "[a-c]"
	ctx := parse(regex)

	expectedLiteralsSet := map[uint8]bool{'a': true, 'b': true, 'c': true}
	expectedTokens := []token{
		{tokenType: bracket, value: expectedLiteralsSet},
	}

	if !reflect.DeepEqual(ctx.tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, got %v", expectedTokens, ctx.tokens)
	}
}

// func TestParseOr(t *testing.T) {
// 	regex := "a|b"
// 	ctx := parse(regex)

// 	expectedTokens := []token{
// 		{tokenType: literal, value: 'a'},
// 		{tokenType: or},
// 		{tokenType: literal, value: 'b'},
// 	}

// 	if !reflect.DeepEqual(ctx.tokens, expectedTokens) {
// 		t.Errorf("Expected tokens %v, got %v", expectedTokens, ctx.tokens)
// 	}
// }

// func TestParseRepeat(t *testing.T) {
// 	regex := "a*"
// 	ctx := parse(regex)

// 	expectedTokens := []token{
// 		{tokenType: literal, value: 'a'},
// 		{tokenType: repeat},
// 	}

// 	if !reflect.DeepEqual(ctx.tokens, expectedTokens) {
// 		t.Errorf("Expected tokens %v, got %v", expectedTokens, ctx.tokens)
// 	}
// }

// func TestParseGroupNested(t *testing.T) {
// 	regex := "(a(bc))"
// 	ctx := parse(regex)

// 	expectedTokens := []token{
// 		{
// 			tokenType: group,
// 			value: []token{
// 				{tokenType: literal, value: 'a'},
// 				{
// 					tokenType: group,
// 					value: []token{
// 						{tokenType: literal, value: 'b'},
// 						{tokenType: literal, value: 'c'},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	if !reflect.DeepEqual(ctx.tokens, expectedTokens) {
// 		t.Errorf("Expected tokens %v, got %v", expectedTokens, ctx.tokens)
// 	}
// }

// func TestParseComplex(t *testing.T) {
// 	regex := "(ab|[a-c]*)"
// 	ctx := parse(regex)

// 	expectedLiteralsSet := map[uint8]bool{'a': true, 'b': true, 'c': true}
// 	expectedTokens := []token{
// 		{
// 			tokenType: group,
// 			value: []token{
// 				{tokenType: literal, value: 'a'},
// 				{tokenType: literal, value: 'b'},
// 				{tokenType: or},
// 				{
// 					tokenType: bracket,
// 					value:     expectedLiteralsSet,
// 				},
// 				{tokenType: repeat},
// 			},
// 		},
// 	}

// 	if !reflect.DeepEqual(ctx.tokens, expectedTokens) {
// 		t.Errorf("Expected tokens %v, got %v", expectedTokens, ctx.tokens)
// 	}
// }
