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
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")

		if len(parts) != 2 {
			fmt.Printf("Invalid input on line %d: %s\n", i+1, lines[i])
			continue
		}

		command := parts[0]
		stepsStr := parts[1]

		steps, err := strconv.Atoi(stepsStr)
		if err != nil {
			fmt.Printf("Invalid steps on line %d: %s\n", i+1, stepsStr)
			continue
		}

		fmt.Printf("Command: %s, Steps: %d\n", command, steps)
	}
}
