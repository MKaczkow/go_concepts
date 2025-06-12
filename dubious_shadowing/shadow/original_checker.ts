function getUnresolvedSymbolForEntityName(
    name: EntityNameOrEntityNameExpression,
) {
    // ...
    let result = unresolvedSymbols.get(path);
    if (!result) {
        unresolvedSymbols.set(
            path,
            result = createSymbol(SymbolFlags.TypeAlias, text),
        );
        result.parent = parentSymbol;
        result.links.declaredType = unresolvedType;
    }
    return result;
}