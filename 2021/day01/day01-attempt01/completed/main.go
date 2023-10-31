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

	var count int = 0

	for i := 0; i < len(lines); i++ {

		if i == 0 {
			continue
		}

		curLine, _ := strconv.Atoi(lines[i])
		prevLine, _ := strconv.Atoi(lines[i-1])

		if i > 0 && curLine > prevLine {
			count = count + 1
			fmt.Println("count increased to:", count)
		}
	}

	fmt.Print(`count is `, count)

}

//////

// func main_old() {
// 	lines := strings.Split(strings.TrimSpace(input), "\n")

// 	var count int = 0

// 	for i := 0; i < len(lines); i++ {
// 		if i > 0 && lines[i] > lines[i-1] {
// 			count++
// 			fmt.Println("count increased to:", count)
// 		}
// 	}

// 	fmt.Print("count is", count)

// for i, line := range lines {
// 	fmt.Println(i, line)

// 	// skip first run
// 	// if i == 0 {
// 	// 	continue
// 	// }

// 	if i > 0 && lines[i] > lines[i-1] {
// 		count++
// 		fmt.Print("count increased to:", count)
// 	}
// }

// fmt.Print("number of times depth has increased:", count)
// }
