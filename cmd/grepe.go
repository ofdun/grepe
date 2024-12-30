package cmd

import (
	"github.com/spf13/cobra"
	myio "grepe/internal/io"
	"grepe/internal/parsers"
	"io"
	"os"
	re "regexp"
)

const (
	MinArgs             = 1
	MaxArgs             = 2
	ArgsLenWithFilename = 2
)

var rootCmd = &cobra.Command{
	Use:   "grepe",
	Short: "grepe is short for grep-extended",
	Long:  `grepe utility searches any given input files, selecting lines that match one or more patterns.`,
	Args:  cobra.RangeArgs(MinArgs, MaxArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		var reader io.Reader
		var err error
		if reader, err = parseFileArgument(args); err != nil {
			return err
		}

		pattern, err := re.Compile(args[0])
		if err != nil {
			return err
		}

		matchingRows, indexes, err := parsers.GetMatchIndexesArray(reader, pattern)
		if err != nil {
			return err
		}

		myio.PrintPatternMatchesColorful(matchingRows, indexes)

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
