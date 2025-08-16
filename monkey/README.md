# Go Concepts
Repo for basic tutorial-based Golang study  

---

# monkey
... programming language

## run
* run REPL -> `go run main.go`
* test module -> `go test ./interpreter/lexer`

## adding new type checklist

### for interpreter
* [ ] add new token type (in `token/token.go`) to convert stream of characters into stream of tokens
	* [ ] define token type 
	* [ ] add branch in `NextToken()` (in `lexer/lexer.go`) function, calling new function 
	* [ ] add function to actually convert characters into tokens of given type
* [ ] add parsing logic to convert stream of tokens into AST (Abstract Syntax Tree)
	* [ ] define node (in `ast/ast.go`)
	* [ ] register prefix or infix parsing function (in `parser/parser.go`, function `New()`)
	* [ ] add parsing logic as separate function(in `parser/parser.go`)
* [ ] add evaluation logic to convert AST into output
	* [ ] define new type object (in `object/object.go`)
	* [ ] add evaluation logic (in `evaluator/evaluator.go`, function `Eval()`)

### for compiler
* [ ] reuse token type from `token.go`
* [ ] add case-branch in `compiler.go` (`Compile` method)
* [ ] add case-branch in `vm.go` (`Run` method)

## notes

### basics
* `compiler` -> parses source code and produces something (output in machine language)
* `interpreter` -> parses source code and doesn't produce antything
	* really simple, don't even bother parsing -> brainfuck
	* tree-walking -> produce AST and evaluate it
	* more complex, like JIT -> get internal representation and then execute on-the-flight
	* `lexer`
	* `parser`
	* tree representation
	* `evaluator`

### lexer
* source code -> tokens
* call `NextToken()` until finished (EOF??)

### parser
* produces `AST` (Abstract Syntax Tree)
* it's *abstract* because not all characters from source code go into the tree
* `CFG` (Context Free Grammar) -> set of rules which describes how to form a correct sentence in given language
* `BNF` (Backus-Naur Form) and `EBNF` (Extended Backus-Naur Form) are notations to represent CFG
* you can generate parser automatically, using definitions above
* `top-down` parsing or `bottom-up` parsing strategies exist, the one written in this exercise will be `top down operator precedence` or `Pratt parser`
* `expression` (produces value) != `statement` (doesn't produce value)
* for example (very basic) this piece of code:  
```monkey
let x = 5;
```  
is represented as:  
![monkey-let-ast](./img/monkey-interpreter-01.png)
* `expression parsing` is likely the most complicated
* nearly all parser share *assertions functions*, which are used to enforce correct order of tokens, so that they make sense, e.g.
```go
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}
```  
* in `monkey` there are only 2 statements: `let` and `return`
* ... but much more expressions
	- `prefix` (e.g. `-5`)
	- `infix` (e.g. `5 + 5`)
	- `postfix` (e.g. `5++`)
	- `call expression` (e.g. `add(1, 2)`)
	- ... 
* `parseExpression` method is recursive, because expressions can be nested, which is the clou of the whole parser idea (*recursive descent parser*)
* `Top Down Operator Precedence` by V.Pratt [link](https://tdop.github.io/), but with differences:
	* `nuds` (null denotation) -> `prefixParseFns`
	* `leds` (left denotation) -> `infixParseFns`
* challenge is, to nest the nodes in AST correctly, so that 
```math
1 + 2 + 3
```  
which *mathematically* is
```math
((1 + 2) + 3)
```  
is represented as:   
![monkey-math-ast](./img/monkey-interpreter-02.png)  
* the higher the precedence, the deeper in the tree the node ends up

### evaluator
* `evaluator` is the part of the interpreter which takes the `AST` and produces the output
* we use `tree-walking interpreter`, which just traverses the tree and evaluates the nodes on the flight
* ... so, we don't use `bytecode` as an `intermediate representation (IR)`
* bytecode is neither native `machine code` nor `assembly language code`
* ... it' instead interpreted by `virtual machine`, which is a part of the interpreter in this case
* `JIT (Just-In-Time)` compiler is a hybrid of `interpreter` and `compiler`, i.e. it receives `bytecode` in virtual machine, which then compiles it to machine code, but only when needed (on the fly), hence the name
* `native types` vs `pointer` as objects representation - design choice
* `self-evaluating expressions` are those which evaluate to themselves, e.g. `5` or `true`, thus being the easiest to implement
* lots of expression in monkey are treated as `prefix expressions`, because they are the easiest to evaluate
* conditionals are built, so that branch is evaluated when expression is `truthy`, which means not null and not false
* implementing internal error-handling is kinda similar to handling return statements, in the sense, that they both stop further evaluation
* `environment` is used to keep track of values of variables, which are stored in `hash map`
* function also carry their own `environment`, which allows using `closures`
* extendint `environment` is done by creating new `environment` with reference to the old one (weird, but alows each `function` to have its own `scope`)

### repl
* `REPL` - `Read-Eval-Print Loop`
* sometimes - 'console', 'interactive mode', etc.

### built-in functions
* design choice - should they be evalueated in top-level environment or in their own environment? (`object.Environment`)
* `index operator` is treated like infix operator, with `someArray` as left operand and `index` as right operand, e.g. `myArray[0]` means `myArray` is left operand and `0` is right operand
* problem with hashes keys (being treated as distinct objects) is solved in `HashKey()` method, which also preserves `O(n)` complexity and insted keeps `O(1)` complexity for accessing hash keys
* any object implementing `Hashable` interface can be used as a key in a hash map

### compiler
* goal is turning `source code` into `bytecode` and then use VM to execute it

### macros
* 2 types of macros:
    - `text-substitution`
    - `syntactic macros` 
* in `Monkey` there will be macros, based on Elixir macros
* macros sit between `parser` and `evaluator`, i.e. they modify `AST` before it's evaluated
* there are 2 stages of `macro expansion`:
    - extracting them from AST
    - evaluating them

### further (possible) exercises
- [ ] compiler
	- [x] chapter 1
	- [x] chapter 2
	- [x] chapter 3
	- [x] chapter 4 (conditionals, expressions)
	- [x] chapter 5 (bindings - `let` and `=`)
	- [x] chapter 6 (3 more data types - arrays, strings, hashes)
	- [x] chapter 7 (functions)
    	- [x] allocate space for local bindings on the stack (i.e. *increase vm.sp* without pushing anything)
    	- [x] implement `OpSetLocal` and `OpGetLocal` instruction in the VM
    	- [x] `arguments` are actually very similar to local bindings, only difference being - they are created implicitly by the compiler and VM
	- [ ] chapter 8
    	- [x] move builtins to `object` package, which is the *right thing to do* design-wise
    	- [ ] add separate scope for built-in functions, as they are neither global nor local scope
	- [ ] chapter 9
	- [ ] chapter 10
	- [x] CI for monkey compiler
	- [ ] stack visualization (probably some simple function in `vm.go` file)
- [ ] misc
    - [ ] add CLI tool (with powerful name, like 'chimp' or 'kong') to trigger compilation (no REPL)
    - [ ] enable compiling from files (like .monkey or smth)
    - [ ] enable multi-files compilation
    - [ ] add stack visualization tool / script
    - [ ] add graceful shutdown to REPL (like `exit` command or `Ctrl+C`)
	- [ ] fully support Unicode in Monkey (as of now, it's only ASCII)
	- [ ] read through Wren language [code](https://github.com/wren-lang/wren)
    - [ ] add nested macros support to monkey interpreter
- [x] interpreter
	- [x] 'talk-me-through' for monkey interpreter chapter 3
	- [x] 'talk-me-through' for monkey interpreter chapter 4
	- [x] check for tests (are they complete?) in array built-in functions
	- [x] find and fix bug in precedence parsing (or tests)
	- [x] chapter 4
	- [x] CI for monkey interpreter
- [x] macros
    - [x] what are macros?
    - [x] add macros to monkey interpreter
      - [x] quote / unquote functions
      - [x] add `MACRO` keyword to lexer
      - [x] add `macro` function to parser
	  - [x] add `macro` function to evaluator
