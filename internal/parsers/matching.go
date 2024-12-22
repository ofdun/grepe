package parsers

import (
	"bufio"
	"io"
	"strings"
)

func getMatchIndexes(s string, pattern string) *[]int {
	var indexes []int
	start := 0

	for {
		index := strings.Index(s[start:], pattern)
		if index == -1 {
			break
		}

		absoluteIndex := start + index
		indexes = append(indexes, absoluteIndex)
		start = absoluteIndex + 1
	}

	return &indexes
}

func GetMatchIndexesArray(strs *[]string, pattern string) (*[][]int, error) {
	indexes := make([][]int, len(*strs))

	for i, str := range *strs {
		indexes[i] = append(indexes[i], *getMatchIndexes(str, pattern)...)
	}

	return &indexes, nil
}

func FindInText(reader io.Reader, pattern string) (*[]string, error) {
	slice := make([]string, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, pattern) {
			slice = append(slice, line)
		}
	}

	return &slice, nil
}
