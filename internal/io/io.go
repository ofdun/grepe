package io

import "fmt"

func PrintPatternMatchesColorful(indexes *[][]int, strs *[]string, patternLength int) {
	green := "\033[32m"
	reset := "\033[0m"

	for i, str := range *strs {
		prev := 0
		for _, val := range (*indexes)[i] {
			fmt.Print(str[prev:val] + green + str[val:(val+patternLength)] + reset)
			prev = val + patternLength
		}
		fmt.Println(str[prev:])
	}
}
