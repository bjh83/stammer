package main

import(
	compiler "../regex_compiler"
	"fmt"
	"os"
)

func main(regex string) {
	instructions := compiler.Compile(regex)
	file, err = os.Create("test_output.txt")
	if err != nil {
		fmt.Println("ERROR: file could not be created")
	}
	for index := 0; index < len(instructions); index++ {
		switch instructions[index].OpCode {
		case Start:
			fmt.Fprintf(file, "START\n")
			break
		case Char:
			fmt.Fprintf(file, "CHAR\t", rune(instructions[index].Line1), "\n")
			break
		case Jump:
			fmt.Fprintf(file, "JMP\t", instructions[index].Line1, "\n")
			break
		case Split:
			fmt.Fprintf(file, "SPLIT\t", instructions[index].Line1,
			instructions[index].Line2, "\n")
			break
		case Match:
			fmt.Fprintf(file, "MATCH")
			break
		}
	}
	fmt.Println("Compilation complete")
}

