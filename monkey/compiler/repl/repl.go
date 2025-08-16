package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/compiler/compiler"
	"monkey/compiler/lexer"
	"monkey/compiler/object"
	"monkey/compiler/parser"
	"monkey/compiler/vm"
)

const PROMPT = ">> "

const MONKEY_FACE = `
     __,__  
    .--.  .-"     "-.  .--.  
   / .. \/  .-. .-.  \/ .. \ 
  | |  '|  /   Y   \  |'  | |
  | \   \ \ 0 | 0 / /   / |  
   \ '- ,\.-"""""""-./, -' /  
    ''-'  /_   ^ ^   _\  '-''  
         | \._     _./ |  
         \ \  '~'~'  / /  
          '._ '-=-' _.'  
             '-----'   
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalsSize)
	symbolTable := compiler.NewSymbolTable()
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		comp := compiler.NewWithState(symbolTable, constants)
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
			continue
		}

		code := comp.Bytecode()
		constants = code.Constants

		machine := vm.NewWithGlobalsStore(code, globals)
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
			continue
		}

		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
