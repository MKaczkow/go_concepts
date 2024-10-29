package parser

import (
	"fmt"
	"strings"
)

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

/*
helper functions for debugging, should be used together, along with `defer`
example usage:
(assume you have function `parseExpressionStatement` in `parser.go`)
```golang
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	//some logic
	return stmt
}
```
just add `trace` and `untrace` to the function:
```golang
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	defer untrace(trace("parseExpressionStatement"))
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	//some logic
	return stmt
}
````
*/

func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}
