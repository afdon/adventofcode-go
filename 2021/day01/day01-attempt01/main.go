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

	// var current [3]int
	// var next [3]int

	var count int = 0

	for i := 0; i < len(lines)-3; i++ {

		var one int = 0
		var two int = 0

		first, _ := strconv.Atoi(lines[i])
		second, _ := strconv.Atoi(lines[i+1])
		third, _ := strconv.Atoi(lines[i+2])
		fourth, _ := strconv.Atoi(lines[i+3])

		one = first + second + third
		two = second + third + fourth

		if two > one {
			count++
		}

		////

		// 	curLine, _ := strconv.Atoi(lines[i])
		// 	prevLine, _ := strconv.Atoi(lines[i-1])

		// 	if i > 0 && curLine > prevLine {
		// 		count = count + 1
		// 		fmt.Println("count increased to:", count)
		// 	}
		// }

		fmt.Print(`count is `, count)

	}

}
