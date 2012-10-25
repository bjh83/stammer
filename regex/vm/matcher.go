package vm

import(
	. "../oplist"
	"./queue"
	"fmt"
)

func ThompsonVM(input string, instructions []Instruct) bool {
	cqueue := queue.New()
	nqueue := queue.New()
	cqueue.Push(0)
	for strCount := 0; strCount <= len(input); strCount++ {
		for cqueue.Len() > 0 {
			instrCount := cqueue.Pop()
			instr := instructions[instrCount]
			fmt.Println("Text pointer: ", strCount, "Instruction count: ", instrCount)
			switch instr.OpCode {
			case Start:
				fmt.Println("START")
				cqueue.Push(1) //This will cause a bootch is Start is not the first instr
				if instrCount != 0 {
					fmt.Println("Total failure")
					return false
				}
				break
			case Char:
				if strCount >= len(input) || input[strCount] != uint8(instr.Line1) {
					fmt.Println("CHAR")
					fmt.Println("Thread died")
					break //Thread dies
				}
				fmt.Println("CHAR ", string(uint8(instr.Line1)), " ", string(input[strCount]))
				nqueue.Push(instrCount + 1)
				break
			case Jump:
				fmt.Println("JUMP")
				cqueue.Push(instr.Line1)
				break
			case Split:
				fmt.Println("SPLIT")
				cqueue.Push(instr.Line1)
				cqueue.Push(instr.Line2)
				break
			case Match:
				fmt.Println("MATCH")
				if strCount == len(input) {
					return true
				}
				break
			}
		}
		fmt.Println("SWITCH")
		cqueue, nqueue = nqueue, cqueue
	}
	return false
}

