package regex_compiler

import(
	"github.com/bjh83/stammer/regex_compiler/lexer"
	"github.com/bjh83/stammer/regex_compiler/parser"
	"github.com/bjh83/stammer/regex_compiler/oplist"
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

