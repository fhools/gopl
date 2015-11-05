package main

import (
	"fmt"
	"os"
	"strings"
)

// Exercise 1.1 Modify program to also print out name of program
func main() {
	fmt.Println(strings.Join(os.Args[:], " "))
}
