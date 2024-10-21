# Go Concepts
Repo for basic tutorial-based Golang study  

---

# monkey
... programming language

## run
* run REPL -> `go run main.go`
* test module -> `go test ./interpreter/lexer`

## adding new type checklist
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

### compiler
* goal is turning `source code` into `bytecode` and then use VM to execute it

### further (possible) exercises
- [ ] interpreter
	- [x] CI for monkey interpreter
	- [ ] 'talk-me-through' for monkey interpreter chapter 3
	- [ ] chapter 4
	- [ ] find and fix bug in precedence parsing (or tests)
	- [ ] check for tests (are they complete?) in array built-in functions
- [ ] compiler
	- [ ] chapter 1
	- [ ] CI for monkey compiler
- [ ] misc [ ] 
	- [ ] fully support Unicode in Monkey (as of now, it's only ASCII)
	- [ ] read through Wren language [code](https://github.com/wren-lang/wren)
