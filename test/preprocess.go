package main

import(
	. "../regex/preprocessor"
	"fmt"
	"os"
)

func main() {
	fmt.Println(PreProcess(os.Args[1]))
}

