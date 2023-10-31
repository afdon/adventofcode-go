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

	// var depth int = 0
	// var position int = 0

	fmt.Println(lines)

	for i := 0; i < len(lines); i++ {

		// lines[i]

		// steps, _ := strconv.Atoi(lines[i][len(lines) - 1])

		// 	dir := dir(lines[i])
		// 	steps := steps(lines[i])

		// 	fmt.Print("dir", dir)
		// 	fmt.Print("steps", steps)

		// 	if dir == 0 {
		// 		position += steps
		// 	}

		// 	if dir == 1 {
		// 		depth += steps
		// 	}

		// 	if dir == -1 {
		// 		depth -= steps
		// 	}
	}

	// fmt.Print(`position is `, position)
	// fmt.Print(`depth is `, depth)
}

/////

// func steps(line string) int {
// 	steps, _ := strconv.Atoi(strings.Split(line, " ")[1])

// 	fmt.Print("steps", steps)

// 	return steps
// }

// func dir(line string) int {

// 	fmt.Print("line", line)

// 	dir := strings.Split(line, " ")[0]

// 	fmt.Print("dir", dir)

// 	if dir == "up" {
// 		fmt.Print("U")
// 		return -1
// 	}
// 	if dir == "forward" {
// 		return 0
// 	}
// 	if dir == "down" {
// 		return 1
// 	}

// 	fmt.Print("couldn't parse direction")
// 	return 0
// }
