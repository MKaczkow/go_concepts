# Go Concepts
Repo for basic tutorial-based Golang study  

---

# monkey
... programming language

## run
* just `make`

## notes
* compiler starts with code from `interpreter`

### basics
* `compiler` -> parses source code and produces something (output in machine language)
* `interpreter` -> parses source code and doesn't produce antything
* `virtual machine` -> *machine built in software*, executing `fetch-decode-execute cycle`
  * `stack machine` -> easier to build, but need to execute lot's of instructions
  * `register machine` -> harder to build
* `little endian` -> least significant byte is stored first
* `big endian` -> most significant byte is stored first

### functions
* notoriously hard, because:
  * first-class citizens => able to pass them around
  * return => how and where?
* curious thing about function literals is that the value they produce doesn't change
* no need for new opcode for `function literals` => they are represented as `object.CompiledFunction`
* ... but we do need new opcode for function call => to tell VM *execute the function on the top of the stack* (`calling convention`) => but how to tell VM *this is the end of function, return*
* `stack of compilation scopes`