package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"grepe/internal/config"
	myio "grepe/internal/io"
	"grepe/internal/parsers"
	"io"
	"os"
	re "regexp"
)

const (
	countDescription        = "Suppress normal output; instead print a count of matching lines."
	ignoreCaseDescription   = "Ignore case distinctions in patterns and input data."
	noIgnoreCaseDescription = "Do not ignore case distinctions in patterns and input data. This is the default. This option is useful for passing to shell scripts that already use -i, in order to cancel its effects."
	invertMatchDescription  = "Invert the sense of matching, to select non-matching lines. (-v is specified by POSIX)"
	wordsOnlyDescription    = "Select only those lines containing matches that form whole words. This option has no effect if -x is also specified."
	fullLineDescription     = "Select only those matches that exactly match the whole line. (-x is specified by POSIX.)"
	colorDescription        = "Color in which matches are highlighted. According to ECMA-48 standard the next values are expected: black, red, green, yellow, blue, magenta, cyan, white, black-background, red-background, green-background, yellow-background, blue-background, magenta-background, cyan-background, white-background. The default color is green"
	maxLinesDescription     = "Stop after the first num selected lines. If num is zero, grepe stops right away without reading input. A negative num is treated as infinity and grepe does not stop; this is the default."
	onlyMatchingDescription = "Print only the matched non-empty parts of matching lines, with each such part on a separate output line."
	quietDescription        = "Quiet; do not write anything to standard output. Exit with zero status if any match is found"
)

const (
	MinArgs             = 1
	MaxArgs             = 2
	ArgsLenWithFilename = 2
	Version             = "0.1"
)

var (
	globalConfig *config.Config
)

var rootCmd = &cobra.Command{
	Use:     "grepe",
	Short:   "grepe is short for grep-extended",
	Long:    `grepe utility searches any given input files, selecting lines that match one or more patterns.`,
	Args:    cobra.RangeArgs(MinArgs, MaxArgs),
	Version: Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.AdaptConfig(cmd.Flags(), globalConfig, args[0]); err != nil {
			return err
		}

		var reader io.Reader
		var err error
		if reader, err = parseFileArgument(args); err != nil {
			return err
		}

		pattern, err := re.Compile(globalConfig.Pattern)
		if err != nil {
			return err
		}

		matchingRows, indexes, err := parsers.GetMatchIndexesArray(reader, pattern, globalConfig)
		if err != nil {
			return err
		}

		if !globalConfig.Quiet {
			if globalConfig.Count {
				fmt.Println(len(*matchingRows))
			} else {
				myio.PrintPatternMatchesColorful(matchingRows, indexes, globalConfig.Color, globalConfig.OnlyMatching)
			}
		} else if len(*matchingRows) == 0 {
			os.Exit(2) // not found anything while being on quiet mode
		}

		return nil
	},
}

// TODO -e and -f

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

func setupFlags(cmd *cobra.Command) *config.Config {
	cfg := &config.Config{}

	cmd.PersistentFlags().BoolVarP(&cfg.Count, "count", "c", false, countDescription)
	cmd.PersistentFlags().BoolVarP(&cfg.IgnoreCase, "ignore-case", "i", false, ignoreCaseDescription)
	cmd.PersistentFlags().BoolVarP(&cfg.Inverted, "invert-match", "v", false, invertMatchDescription)
	cmd.PersistentFlags().Bool("no-ignore-case", false, noIgnoreCaseDescription)
	cmd.PersistentFlags().BoolVarP(&cfg.WordsOnly, "word-regexp", "w", false, wordsOnlyDescription)
	cmd.PersistentFlags().BoolVarP(&cfg.FullLine, "line-regexp", "x", false, fullLineDescription)
	cmd.PersistentFlags().BoolVarP(&cfg.OnlyMatching, "only-matching", "o", false, onlyMatchingDescription)
	cmd.PersistentFlags().IntVarP(&cfg.MaxLines, "max-count", "m", -1, maxLinesDescription)
	cmd.PersistentFlags().StringVar(&cfg.Color, "color", "green", colorDescription)
	cmd.PersistentFlags().BoolVarP(&cfg.Quiet, "quiet", "q", false, quietDescription)

	cmd.PersistentFlags().BoolVar(&cfg.Quiet, "silent", false, quietDescription)
	cmd.PersistentFlags().StringVar(&cfg.Color, "colour", "green", "Alias for --color. Overrides it")

	return cfg
}

func Execute() {
	globalConfig = setupFlags(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
