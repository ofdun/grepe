package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	myio "grepe/internal/io"
	"grepe/internal/parsers"
	"io"
	"os"
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
		result, err := parsers.FindInText(reader, pattern)
		if err != nil {
			return err
		}

		indexes, err := parsers.GetMatchIndexesArray(result, pattern)
		if err != nil {
			return err
		}

		myio.PrintPatternMatchesColorful(indexes, result, len(pattern))

		return nil
	},
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
