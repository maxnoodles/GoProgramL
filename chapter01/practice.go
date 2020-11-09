package main

import (
	"fmt"
	"os"
	"strings"
)

func echo() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echo1() {
	var s, sep string
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}

func main() {
	echo2()
}
