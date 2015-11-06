package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// map of line content to  map[filename]counts. counts isn't needed
	dupfiles := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts, dupfiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts, dupfiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for filename, _ := range dupfiles[line] {
				fmt.Printf("%s\n", filename)
			}
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]int, dupfiles map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		// create initial map if not created yet
		if dupfiles[input.Text()] == nil {
			dupfiles[input.Text()] = make(map[string]int)
		}
		dupfiles[input.Text()][filename]++
	}
}

