package parsers

import (
	"bufio"
	"grepe/internal/config"
	"io"
	re "regexp"
)

func handleInvertedIndexesCase(indexes *[][]int, slice *[][][]int, text *[]string, line *string) {
	if indexes == nil {
		*text = append(*text, *line)

		fullLineIndexes := make([][]int, 1)
		fullLineIndexes[0] = make([]int, 2)

		fullLineIndexes[0][0] = 0
		fullLineIndexes[0][1] = len(*line)

		*slice = append(*slice, fullLineIndexes)
	} else {
		if len(*indexes) == 1 && (*indexes)[0][0] == 0 && (*indexes)[0][1] == len(*line) {
			return
		}

		*indexes = invertIndexes(*indexes, len(*line))
		*text = append(*text, *line)
		*slice = append(*slice, *indexes)
	}
}

func GetMatchIndexesArray(reader io.Reader, pattern *re.Regexp, cfg *config.Config) (*[]string, *[][][]int, error) {
	slice := make([][][]int, 0)
	text := make([]string, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if cfg.MaxLines == 0 {
			break
		} else if cfg.MaxLines > 0 {
			cfg.MaxLines--
		}

		line := scanner.Text()
		indexes := pattern.FindAllStringIndex(line, -1)

		if !cfg.Inverted {
			if indexes != nil {
				text = append(text, line)
				slice = append(slice, indexes)
			}
		} else {
			handleInvertedIndexesCase(&indexes, &slice, &text, &line)
		}
	}

	return &text, &slice, nil
}

func invertIndexes(indexes [][]int, lengthOfLine int) [][]int {
	res := make([][]int, 0, len(indexes)+1)
	prev := 0
	for _, cur := range indexes {
		from := cur[0]
		to := cur[1]

		if from != prev {
			newInterval := []int{prev, from}
			res = append(res, newInterval)
		}

		prev = to
	}

	if prev != lengthOfLine {
		newInterval := []int{prev, lengthOfLine}
		res = append(res, newInterval)
	}

	return res
}
