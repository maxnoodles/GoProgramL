package main

import (
	"bufio"
	"fmt"
	"os"
)

type line struct {
	Filename string
	String   string
}

func main() {
	counts := make(map[line]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s%n\n", line.Filename, line.String)
		}
	}
}

func countLines(f *os.File, counts map[line]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[line{f.Name(), input.Text()}]++
	}
}
