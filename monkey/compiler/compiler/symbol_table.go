package compiler

type SymbolScope string

const (
	GlobalScope SymbolScope = "GLOBAL"
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

// Symbol table is a data structure used in interpreters and compilers
// to associate identifiers with information (symbols).
// Common 2 methods on a symbol table are:
// - Define()
// - Resolve()
//
// We will use it to associate identifiers with:
// - scope
// - unique number
//
// Reference:
// https://en.wikipedia.org/wiki/Symbol_table
type SymbolTable struct {
	store          map[string]Symbol
	numDefinitions int
}

func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Scope: GlobalScope, Index: s.numDefinitions}
	s.store[name] = symbol
	s.numDefinitions++
	return symbol
}

func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	return obj, ok
}
