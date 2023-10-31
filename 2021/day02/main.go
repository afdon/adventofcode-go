package main

import (
	_ "embed"
	"fmt"
	"strings"
)

// go:embed input.txt
var input string

func main() {
	fmt.Println(input)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	var count int = 0

	for i := 0; i < len(lines)-3; i++ {

		fmt.Print(`count is `, count)

	}

}
