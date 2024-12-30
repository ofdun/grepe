package parsers

import (
	"bufio"
	"io"
	re "regexp"
)

func GetMatchIndexesArray(reader io.Reader, pattern *re.Regexp) (*[]string, *[][][]int, error) {
	slice := make([][][]int, 0)
	text := make([]string, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		indexes := pattern.FindAllStringIndex(line, -1)

		if indexes != nil {
			text = append(text, line)
			slice = append(slice, indexes)
		}
	}

	return &text, &slice, nil
}
