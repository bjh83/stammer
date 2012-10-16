package compiler

import(
	"fmt"
	"parsetree"
)

const(
	Pipe = 256
	Star
	Plus
	Ques
)

func lex(regex string) out string {
	nextEscape := false
	escaped := false
	for i := 1; i < len(regex); i++ {
		switch regex[i] {
		case '\\':
			if escaped {
				out += '\\'
				escaped = false
			} else {
				nextEscape = true
			}
		case '|':
			if escaped {
				out += '|'
				escaped = false
			} else {
				out += Pipe
			}
		case '*':
			if escaped {
				out += '*'
				escaped = false
			} else {
				out += Star
			}
		case '+':
			if escaped {
				out += '+'
				escaped = false
			} else {
				out += Plus
			}
		case '?':
			if escaped {
				out += '?'
				escaped = false
			} else {
				out += Ques
			}
		default:
			out += regex[i]
		}
		if escaped {
			//ERROR: Nothing consumed the escape
			fmt.Println("The escape before character: ", regex[i], " index: ", i, "was not consumed!")
		}
		escaped = nextEscape
		nextEscape = false
	}
}

