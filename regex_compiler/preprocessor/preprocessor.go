package preprocessor

import(
	. "../lexer"
	. "container/list"
)

const(
	ClassOpen = iota + Buffer
	ClassClose = iota + Buffer
	Delete = iota + Buffer
)

func PreProcess(regex string) string {
	lexed := lex(regex)
	return expand(lexed)
}

func lex(regex string) []int {
	nextEscape := false
	escaped := false
	var toDelete *int
	out := make([]int, len(regex))
	out_index := 0
	for reg_index := 0; reg_index < len(regex); reg_index++ {
		escaped = nextEscaped
		nextEscaped = false
		switch regex[reg_index] {
		case '\\':
			if escaped {
				escaped = false
				out[out_index] = '\\'
				out_index++
				toDelete = nil
			} else {
				nextEscaped = true
				out[out_index] = '\\'
				toDelete = &out[out_index]
				out_index++
			}
			break
		case '[':
			if escaped {
				escaped = false
				out[out_index] = '['
				*toDelete = Delete
				out_index++
				toDelete = nil
			} else {
				out[out_index] = ClassOpen
				out_index++
			}
			break
		case ']':
			if escaped {
				escaped = false
				out[out_index] = ']'
				*toDelete = Delete
				out_index++
				to_Delete = nil
			} else {
				out[out_index] = ClassClose
				out_index++
			}
			break
		}
	}
	return out
}

func expandClass(regex []int) string {
	list := New()
	list.Init()
	nextEscape := false
	escaped := false
	var toDelete *int
	for reg_index := 0; reg_index < len(regex); reg_index++ {
		escaped = nextEscaped
		nextEscaped = false
		switch regex[reg_index] {
		case '\\':
			if escaped {
				escaped = false
				list.PushBack(int('\\'))
				toDelete = nil
			} else {
				nextEscaped = true
				toDelete = &int('\\')
				list.PushBack(*toDelete)
			}
			break
		case '-':
			if escaped {
				escaped = false
				list.PushBack(int('-'))
				*toDelete = Delete
			} else {
				if reg_index > 1 && reg_index < len(regex) - 1 {
					start := regex[reg_index - 1]
					end := regex[reg_index + 1]
					for char := uint8(start + 1); char < end; char++ {
						list.PushBack(int(char))
					}
				} else {
					//ERROR
					fmt.Println("invalid character class")
				}
			}
			break
		default:
			list.PushBack(regex[reg_index])
			break
		}
	}
	for e := list.Front(); e != nil; e = e.Next() { //Cleans out deletions
		if int(e.Value) == Delete {
			list.Remove(e)
		}
	}
	out := string(list.Remove(list.Front()))
	for e := list.Front(); e != nil; e = e.Next() {
		out += string('|')
		out += string(e.Value)
	}
	return out
}

func expand(regex []int) string {
	var out string
	for index := 0; index < len(regex); index++ {
		if regex[index] == ClassOpen {
			start := index + 1
			for ; index != ClassClose; index++ {}
			end := index
			out += "("
			out += expandClass(regex[start:end])
			out += ")"
		} else {
			out += string(regex[index])
		}
	}
	return out
}

