package evaluator

import (
	"monkey/interpreter/ast"
	"monkey/interpreter/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
