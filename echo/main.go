package main

import (
	"fmt"
	"os"
)

func main() {
	var output, sep string
	for i := 1; i < len(os.Args); i++ {
		output += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(output)
}
