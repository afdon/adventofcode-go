package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(input)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for i := 0; i < len(lines); i++ {

		l, _ := strconv.Atoi(lines[i])

		d := lines[i][len(lines[i])-1]
		s := lines[i][0]

		fmt.Println(l, d, s)

	}

}
