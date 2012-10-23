package regex_compiler

import(
	"./lexer"
	"./parser"
	"./oplist"
	"fmt"
)

func Compile(regex string) []oplist.Instruct {
	lexed := lexer.Lex(regex)
	success, parseTree := parser.Parse(lexed)
	if !success {
		fmt.Println("Parsing Failed")
		return nil
	}
	return parseTree.Generate().ToArray()
}

