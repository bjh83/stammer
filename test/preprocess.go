package main

import(
	. "../regex_compiler/preprocessor"
	"fmt"
	"os"
)

func main() {
	fmt.Println(PreProcess(os.Args[1]))
}

