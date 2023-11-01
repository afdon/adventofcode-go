package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	raw := strings.Split(strings.TrimSpace(input), "\n")

	lines := raw // copy the array by value, not reference (&raw)

	rp := &raw
	lp := &lines
	fmt.Printf("%v, %v", &rp, &lp)

	// oxygen := oxygen(lines)
	// co2 := co2(lines)

	// oxygenInt, _ := strconv.Atoi(string(oxygen))
	// co2Int, _ := strconv.Atoi(string(co2))

	// power := binaryToDecimal(oxygenInt) * binaryToDecimal(co2Int)

	// fmt.Println(oxygenInt, co2Int, power)
}

func tally(lines []string) []int {
	tally := make([]int, len(lines[0]))

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
			// keep all the elements with zeroes
			tally[i] = '0'
		} else {
			// keep all the elements with ones
			tally[i] = '1'
		}
	}

	fmt.Println(tally)
	return tally
}

func tallyco2(lines []string) []int {
	tally := make([]int, len(lines[0]))

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
			tally[i] = '1'
		} else {
			tally[i] = '0'
		}
	}

	fmt.Println(tally)
	return tally
}

func oxygen(lines []string, tally []byte) []byte {

	oxygen := make([]byte, len(lines[0]))

	filteredLines := lines

	for i := 0; i < len(tally); i++ {
		for j := 0; j < len(filteredLines); j++ {

			if len(oxygen) == 1 {
				return oxygen
			}

			if lines[j][i] != tally[i] {
				// remove it from array.
				filteredLines = append(filteredLines[:j-1], filteredLines[j+1:]...)
			}
		}
	}
	return oxygen
}

func co2(lines []string, tally []byte) []byte {

	co2 := make([]byte, len(lines[0]))

	filteredLines := lines

	for i := 0; i < len(tally); i++ {
		for j := 0; j < len(filteredLines); j++ {

			if len(co2) == 1 {
				return co2
			}

			if lines[j][i] != tally[i] {
				// remove it from array.
				filteredLines = append(filteredLines[:j-1], filteredLines[j+1:]...)
			}
		}
	}
	return co2
}

func binaryToDecimal(num int) int {
	var remainder int
	index := 0
	dec := 0
	for num != 0 {
		remainder = num % 10
		num = num / 10
		dec = dec + remainder*int(math.Pow(2, float64(index)))
		index++
	}
	return dec
}
