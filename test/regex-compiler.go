package main

import(
	compiler "../regex_compiler"
	. "../regex_compiler/oplist"
	"fmt"
	"os"
)

func main() {
	instructions := compiler.Compile(os.Args[1])
	file, err := os.Create("test_output.txt")
	if err != nil {
		fmt.Println("ERROR: file could not be created")
	}
	for index := 0; index < len(instructions); index++ {
		switch instructions[index].OpCode {
		case Start:
			fmt.Fprintf(file, "START\n")
			break
		case Char:
			fmt.Fprintf(file, "CHAR\t%c\n", instructions[index].Line1)
			break
		case Jump:
			fmt.Fprintf(file, "JMP\t\t%d\n", instructions[index].Line1 + 1)
			break
		case Split:
			fmt.Fprintf(file, "SPLIT\t%d\t%d\n", instructions[index].Line1 + 1,
			instructions[index].Line2 + 1)
			break
		case Match:
			fmt.Fprintf(file, "MATCH")
			break
		}
	}
	fmt.Println("Compilation complete")
}

