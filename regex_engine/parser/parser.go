package parser

import "fmt"

type tokenType uint8

const (
	group           tokenType = iota
	bracket         tokenType = iota
	or              tokenType = iota
	repeat          tokenType = iota
	literal         tokenType = iota
	groupUncaptured tokenType = iota
)

type token struct {
	tokenType tokenType
	value     interface{}
}

type parseContext struct {
	pos    int
	tokens []token
}

func parse(regex string) *parseContext {
	ctx := &parseContext{pos: 0, tokens: []token{}}
	for ctx.pos < len(regex) {
		process(regex, ctx)
		ctx.pos++
	}

	return ctx
}

func process(regex string, ctx *parseContext) {
	ch := regex[ctx.pos]
	switch ch {
	case '(':
		// it's a group
		groupCtx := &parseContext{
			pos:    ctx.pos,
			tokens: []token{},
		}
		parseGroup(regex, groupCtx)
		ctx.tokens = append(ctx.tokens, token{
			tokenType: group,
			value:     groupCtx.tokens,
		})
	case '[':
		// it's a bracket expression
		parseBracket(regex, ctx)
	case '|':
		// it's an OR operator
		parseOr(regex, ctx)
	case '*':
		// it's a repeat operator
		// remaining repeat operators are + and ?, but they can be specified using brackets
		// a* == a{0,}
		// a+ == a{1,}
		// a? == a{0,1}
		parseRepeat(regex, ctx)
	case '{':
		parseRepeatSpecified(regex, ctx)
	default:
		// it's a literal (nothing matched)
		t := token{tokenType: literal, value: ch}
		ctx.tokens = append(ctx.tokens, t)
	}

}

func parseGroup(regex string, ctx *parseContext) {
	ctx.pos++
	for regex[ctx.pos] != ')' {
		process(regex, ctx)
		ctx.pos++
	}
}

func parseBracket(regex string, ctx *parseContext) {
	ctx.pos++
	var literals []string
	for regex[ctx.pos] != ']' {
		ch := regex[ctx.pos]

		if ch == '-' {
			// range indicator
			next := regex[ctx.pos+1]
			prev := literals[len(literals)-1][0]
			literals[len(literals)-1] = fmt.Sprintf("%c%c", prev, next) // <3-2>
			ctx.pos++
		} else {
			literals = append(literals, fmt.Sprintf("%c", ch))
		}

		ctx.pos++
	}

	literalsSet := map[uint8]bool{} // because literals ranges can overlap

	for _, l := range literals {
		for i := l[0]; i <= l[1]; i++ {
			literalsSet[i] = true
		}
	}

	ctx.tokens = append(ctx.tokens, token{
		tokenType: bracket,
		value:     literalsSet,
	})
}

func parseOr(regex string, ctx *parseContext) {
	// TODO: implement
}

func parseRepeat(regex string, ctx *parseContext) {
	// TODO: implement
}

func parseRepeatSpecified(regex string, ctx *parseContext) {
	// TODO: implement
}
