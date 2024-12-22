package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

const (
	MinArgs             = 1
	MaxArgs             = 2
	ArgsLenWithFilename = 2
)

var (
	InvalidArgsQuantityError = errors.New("too many args")
)

var rootCmd = &cobra.Command{
	Use:   "grepe",
	Short: "grepe is short for grep-extended",
	Long: `grepe utility searches any given input files, selecting lines 
				that match one or more patterns.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > MaxArgs || len(args) < MinArgs {
			return InvalidArgsQuantityError
		}

		var reader io.Reader
		var err error
		if reader, err = parseFileArgument(args); err != nil {
			return err
		}

		pattern := args[0]
		result, err := FindInText(reader, pattern)
		if err != nil {
			return err
		}

		indexes, err := getMatchIndexesArray(result, pattern)
		if err != nil {
			return err
		}

		PrintPatternMatchesColorful(indexes, result, len(pattern))

		return nil
	},
}

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

func getMatchIndexesArray(strs *[]string, pattern string) (*[][]int, error) {
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

func parseFileArgument(args []string) (io.Reader, error) {
	var err error
	reader := os.Stdin

	if len(args) == ArgsLenWithFilename {
		if reader, err = os.Open(args[len(args)-1]); err != nil {
			return nil, err
		}
	}

	return reader, nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
