package main

import (	
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	switch {
	case len(os.Args) < 2:
		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			fmt.Println(strings.ToUpper(in.Text()))
		}
		if err := in.Err(); err != nil {
			panic(err)
		}
	default:
		fmt.Println(strings.ToUpper(strings.Join(os.Args[1:], " ")))
	}
}
