# Go Concepts
Repo for basic tutorial-based Golang study  

---

# monkey
... programming language

### run
* test module -> `go test ./interpreter/lexer`

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

### repl
* `REPL` - `Read-Eval-Print Loop`
* sometimes - 'console', 'interactive mode', etc.


### further (possible) exercises
- [ ] fully support Unicode in Monkey (as of now, it's only ASCII)
