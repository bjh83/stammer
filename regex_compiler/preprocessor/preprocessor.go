package preprocessor

import(
	. "../lexer"
	. "container/list"
	"fmt"
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
		escaped = nextEscape
		nextEscape = false
		switch regex[reg_index] {
		case '\\':
			if escaped {
				escaped = false
				out[out_index] = '\\'
				out_index++
				toDelete = nil
			} else {
				nextEscape = true
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
				fmt.Println("ClassOpen at: ", out_index)
			}
			break
		case ']':
			if escaped {
				escaped = false
				out[out_index] = ']'
				*toDelete = Delete
				out_index++
				toDelete = nil
			} else {
				out[out_index] = ClassClose
				out_index++
				fmt.Println("ClassClose at: ", out_index)
			}
			break
		default:
			escaped = false
			out[out_index] = int(regex[reg_index])
			out_index++
		}
	}
	return out
}

func expandClass(regex []int) string {
	list := New()
	list.Init()
	nextEscape := false
	escaped := false
	var toDelete *Element
	for reg_index := 0; reg_index < len(regex); reg_index++ {
		escaped = nextEscape
		nextEscape = false
		switch regex[reg_index] {
		case '\\':
			if escaped {
				escaped = false
				list.PushBack(int('\\'))
				toDelete = nil
			} else {
				nextEscape = true
				toDelete = list.PushBack(int('\\'))
			}
			break
		case '-':
			if escaped {
				escaped = false
				list.PushBack(int('-'))
				toDelete.Value = Delete
			} else {
				if reg_index > 0 && reg_index < len(regex) - 1 {
					start := regex[reg_index - 1]
					end := uint8(regex[reg_index + 1])
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
	for e := list.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == Delete {
			list.Remove(e)
		}
	}
	out := string(list.Remove(list.Front()).(int))
	for e := list.Front(); e != nil; e = e.Next() {
		out += string('|')
		out += string(e.Value.(int))
	}
	return out
}

func expand(regex []int) string {
	var out string
	for index := 0; index < len(regex); index++ {
		if regex[index] == ClassOpen {
			start := index + 1
			for ; regex[index] != ClassClose; index++ {}
			end := index
			fmt.Println("start: ", start, "\tend: ", end)
			out += "("
			out += expandClass(regex[start:end])
			out += ")"
		} else if regex[index] != Delete {
			out += string(regex[index])
		}
	}
	return out
}

