package main

import (
	"fmt"
	"io"
	"strings"
)

type str struct {
	name string
}

func main() {
	var str = strings.NewReader("Insert text here")
	printRead(str)
}

func printRead(str *strings.Reader) {
	var b = make([]byte, 5)
	for {
		n, err := str.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
