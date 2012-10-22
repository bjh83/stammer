package lexer

import("fmt")

const(Buffer = 256)

const(
	Pipe = iota + Buffer
	Star = iota + Buffer
	Plus = iota + Buffer
	Ques = iota + Buffer
)

func Lex(regex string) []int {
	nextEscape := false
	escaped := false
	out := make([]int, len(regex))
	out_index := 0
	for i := 0; i < len(regex); i++ {
		switch regex[i] {
		case '\\':
			if escaped {
				out[out_index] = '|'
				out_index++
				escaped = false
			} else {
				nextEscape = true
			}
		case '|':
			if escaped {
				out[out_index] = '|'
				out_index++
				escaped = false
			} else {
				out[out_index] = Pipe
				out_index++
			}
		case '*':
			if escaped {
				out[out_index] = '*'
				out_index++
				escaped = false
			} else {
				out[out_index] = Star
				out_index++
			}
		case '+':
			if escaped {
				out[out_index] = '+'
				out_index++
				escaped = false
			} else {
				out[out_index] = Plus
				out_index++
			}
		case '?':
			if escaped {
				out[out_index] = '?'
				out_index++
				escaped = false
			} else {
				out[out_index] = Ques
				out_index++
			}
		default:
			out[out_index] = int(regex[i])
			out_index++
		}
		if escaped {
			//ERROR: Nothing consumed the escape
			fmt.Println("The escape before character: ", regex[i], " index: ", i, "was not consumed!")
		}
		escaped = nextEscape
		nextEscape = false
	}
	return out
}

