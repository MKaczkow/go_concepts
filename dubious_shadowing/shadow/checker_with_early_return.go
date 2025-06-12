package shadow

func (c *Checker) getUnresolvedSymbolForEntityName(name *ast.Node) *ast.Symbol {
	// ...
	if result := c.unresolvedSymbols[path]; result != nil {
		return result
	}
	result := c.newSymbol(ast.SymbolFlagsTypeAlias, text)
	c.unresolvedSymbols[path] = result
	result.Parent = parentSymbol
	c.declaredTypeLinks.Get(result).declaredType = c.unresolvedType
	return result
}
