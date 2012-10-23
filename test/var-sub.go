package main

import(
	. "../regex_compiler"
	. "../regex_compiler/preprocessor"
	"fmt"
)

func main() {
	Declare("digit", "[0-9]")
	Declare("letter", "[a-zA-Z]")
	fmt.Println(PreProcess("{digit}+|{letter}"))
	fmt.Println(PreProcess("\\{letter\\}|[_%]"))
}

