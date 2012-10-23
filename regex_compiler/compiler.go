package regex_compiler

import(
	"./lexer"
	"./parser"
	"./oplist"
	"./preprocessor"
	"fmt"
)

func Compile(regex string) []oplist.Instruct {
	regex = preprocessor.PreProcess(regex)
	lexed := lexer.Lex(regex)
	success, parseTree := parser.Parse(lexed)
	if !success {
		fmt.Println("Parsing Failed")
		return nil
	}
	return parseTree.Generate().ToArray()
}

func Declare(name, regex string) {
	Variables[name] = regex
}

