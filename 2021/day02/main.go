package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	// fmt.Println(input)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for i := 0; i < len(lines); i++ {
		fmt.Println(strings.Split(lines[i], " "))
		instruction := strings.Split(lines[i], " ")
		dir := instruction[0]
		steps := instruction[1]

		fmt.Print(dir)

		// raw := strings.Split(lines[i], " ")

		// fmt.Println(raw)

		// l, _ := strconv.Atoi(lines[i])

		// line := lines[i]
		// lineLength := len(line)
		// lastIndex := line[lineLength-1]

		// fmt.Println(line, lineLength, lastIndex)

		// d := lines[i][lastIndex]

		// var s byte = lines[i][0]
		// steps := strings.LastIndexByte(lines[i], 0)

		// fmt.Print(steps)

		// fmt.Println(d, s)

	}

}
