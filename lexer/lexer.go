package lexer

import(
	"../regex"
	"container/list"
	"fmt"
)

const(
	ID = iota
	VAL = iota
	NIL = iota
)

type Token struct {
	Type, Value int
	ID string
}

func Lex(input string) list.List {
	return scan(input, funcArray)
}

func check(input string, regexs []func(string)(int, int)) list.List {
	output := list.New()
	output.Init()
	for regex := range regexs {
		_type, idOrVal := regex(input)
		token := Token{Type: _type}
		if idOrVal == ID {
			token.ID = input
		} else if idOrVal == VAL {
			token.Value = int(input)
		}
		if _type != NULL {
			output.PushBack()
		}
	}
	return output
}

func scan(input string, regexs []func(string)(int, int)) list.List {
	output := list.New()
	output.Init()
	rightIndex := 0
	for leftIndex := 0; leftIndex < len(input); leftIndex = rightIndex {
		lastPass := false
		rightIndex++
		for {
			tokens := check(input[leftIndex : rightIndex], regexs)
			if lastPass {
				if leftIndex >= rightIndex {
					fmt.Println("Invalid character sequence at: ", input[leftIndex : leftIndex + 10])
				}
				output.PushBack(tokens.Front().Value.(Token))
				break
			}
			if tokens.Len() > 0 {
				rightIndex++
			} else {
				rightIndex--
				lastPass = true
			}
		}
	}
	return output
}

