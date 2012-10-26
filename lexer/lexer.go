package lexer

import(
	"../regex"
	"container/list"
	"fmt"
)

const(
	NULL
)

var any string

func setup() {
	for var i uint8 = 0; i < 128; i++ {
		any += string(i)
	}
}

type Token struct {
	Type, ID, Value int
}

func lex(input string, regexs []func(string) Token) list.List {
	var output list.List
	for regex := range regexs {
		token := regex(input)
		if token.Type != NULL {
			output.PushBack(token)
		}
	}
	return output
}

