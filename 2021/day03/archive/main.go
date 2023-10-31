package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// go:embed input.txt
var input string

func main() {

	lines := strings.Split(strings.TrimSpace(input), "\n")

	gamma := gamma(lines)
	epsilon := epsilon(lines)

	gammaInt, _ := strconv.Atoi(string(gamma))
	epsilonInt, _ := strconv.Atoi(string(epsilon))

	power := binaryToDecimal(gammaInt) * binaryToDecimal(epsilonInt)

	fmt.Println(gammaInt, epsilonInt, power)

}

func gamma(lines []string) []byte {

	var gamma []byte

	for i := 0; i < len(lines[0]); i++ {

		ones := 0
		zeroes := 0

		for j := 0; j < len(lines); j++ {
			fmt.Print(lines[j])
			if string(lines[j][i]) == "0" {
				zeroes++
				// fmt.Print(zeroes)
			}
			if string(lines[j][i]) == "1" {
				ones++
				// fmt.Print(ones)
			}
		}

		if zeroes > ones {
			gamma[i] = 0
		}
		if ones > zeroes {
			gamma[i] = 1
		}
	}

	fmt.Println(gamma)
	return gamma
}

func epsilon(lines []string) []byte {

	var epsilon []byte

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
			epsilon[i] = 0
		}
		if ones > zeroes {
			epsilon[i] = 1
		}
	}

	fmt.Println(epsilon)
	return epsilon
}

func binaryToDecimal(num int) int {
	var remainder int
	index := 0
	decimalNum := 0
	for num != 0 {
		remainder = num % 10
		num = num / 10
		decimalNum = decimalNum + remainder*int(math.Pow(2, float64(index)))
		index++
	}
	return decimalNum
}
