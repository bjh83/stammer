package main

import(
	"../regex"
	. "../regex/oplist"
	"fmt"
)

func main() {
	code := regex.Compile("(abc)(de)")
	instructions := code.Instructions
	for index := 0; index < len(instructions); index++ {
		switch instructions[index].OpCode {
		case Start:
			fmt.Println(index, "\tSTART")
			break
		case Char:
			fmt.Println(index, "\tCHAR\t", string(instructions[index].Line1))
			break
		case Jump:
			fmt.Println(index, "\tJMP\t", instructions[index].Line1)
			break
		case Split:
			fmt.Println(index, "\tSPLIT\t", instructions[index].Line1, "\t",
			instructions[index].Line2)
			break
		case Match:
			fmt.Println(index, "\tMATCH")
			break
		}
	}
	fmt.Println("Compilation complete")
	//fmt.Println(code.Match("."))
}

