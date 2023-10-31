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

	gamma := gamma(lines)
	epsilon := epsilon(lines)

	gammaInt, _ := strconv.Atoi(string(gamma))
	epsilonInt, _ := strconv.Atoi(string(epsilon))

	power := gammaInt * epsilonInt

	fmt.Println(gammaInt, epsilonInt, power)
}

func gamma(lines []string) []byte {
	gamma := make([]byte, len(lines[0]))

	for i := 0; i < len(lines[0]); i++ {
		ones := 0
		zeroes := 0

		for j := 0; j < len(lines); j++ {
			if string(lines[j][i]) == "0" {
				zeroes++
			}
			if string(lines[j][i]) == "1" {
				ones++
			}
		}

		if zeroes > ones {
			gamma[i] = '0'
		}
		if ones > zeroes {
			gamma[i] = '1'
		}
	}

	fmt.Println(string(gamma))
	return gamma
}

func epsilon(lines []string) []byte {
	epsilon := make([]byte, len(lines[0]))

	for i := 0; i < len(lines[0]); i++ {
		ones := 0
		zeroes := 0

		for j := 0; j < len(lines); j++ {
			if string(lines[j][i]) == "0" {
				zeroes++
			}
			if string(lines[j][i]) == "1" {
				ones++
			}
		}

		if zeroes > ones {
			epsilon[i] = '0'
		}
		if ones > zeroes {
			epsilon[i] = '1'
		}
	}

	fmt.Println(string(epsilon))
	return epsilon
}
