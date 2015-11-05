package main

import (
	"fmt"
	"os"
)
// Modify echo program to print index and argument one per line
func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
