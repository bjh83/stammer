package main

import(
	"../regex"
	"fmt"
)

func main() {
	fmt.Println(regex.Compile("[abc]").Match("a"))
}

