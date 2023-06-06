package explainer

import (
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
)

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, repl.MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func Start(in io.Reader, out io.Writer) {
	bytes, err := io.ReadAll(in)
	if err != nil {
		fmt.Println(err)
	}
	codes := string(bytes)
	l := lexer.New(codes)
	p := parser.New(l)
	program := p.ParseProgram()
	evaluated := evaluator.Eval(program, object.NewEnvironment())
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
	}
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}
