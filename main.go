package main

import (
	"fmt"
	"strconv"
)

func main() {

	n := int64(529)
	fmt.Println(Solution(n))

}

//Solution solves the problem.
func Solution(n int64) int {
	strBin := strconv.FormatInt(n, 2)
	longest := 0
	current := 0
	for i := 0; i < len(strBin); i++ {
		bit := string(strBin[i])

		if bit == "0" {
			current++
		}

		if bit == "1" {
			if current > longest {
				longest = current
			}
			current = 0
		}

	}
	return longest
}
