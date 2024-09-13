package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/interpreter/lexer"
	"monkey/interpreter/parser"
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

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
