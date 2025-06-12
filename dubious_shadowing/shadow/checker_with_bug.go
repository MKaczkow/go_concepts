package shadow

func (c *Checker) getUnresolvedSymbolForEntityName(name *ast.Node) *ast.Symbol {
	// ...
	result := c.unresolvedSymbols[path]
	if result == nil {
		result := c.newSymbol(ast.SymbolFlagsTypeAlias, text)
		// ':=' creates a new symbol within the scope of if statement block
		// thus this function will always return nil
		c.unresolvedSymbols[path] = result
		result.Parent = parentSymbol
		c.declaredTypeLinks.Get(result).declaredType = c.unresolvedType
	}
	return result
}
