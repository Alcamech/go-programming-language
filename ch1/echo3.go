package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i, args := range os.Args[1:] {
		fmt.Println("i", i, "args", args)
	}
	fmt.Println(strings.Join(os.Args[1:], " "))
}
