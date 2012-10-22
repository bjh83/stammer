package regex_compiler

import(
	"./lexer"
	"./parser"
	"./oplist"
	"fmt"
)

func Compile(regex string) []Instruct {
	regex = lexer.Lex(regex)
	success, parseTree = parser.Parse(regex)
	if !success {
		fmt.Println("Parsing Failed")
		return nil
	}
	return parseTree.Generate().ToArray()
}

