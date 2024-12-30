package config

import (
	"errors"
	"github.com/spf13/pflag"
	myio "grepe/internal/io"
)

var (
	ErrorUnknownColor = errors.New("unknown color")
)

// TODO exit immediately with first match using quiet flag

type Config struct {
	Pattern      string
	Color        string
	Count        bool
	IgnoreCase   bool
	Inverted     bool
	WordsOnly    bool
	FullLine     bool
	OnlyMatching bool
	Quiet        bool
	MaxLines     int
}

func AdaptConfig(flags *pflag.FlagSet, cfg *Config, pattern string) error {
	noIgnoreCase, err := flags.GetBool("no-ignore-case")
	if err != nil {
		return err
	}

	if noIgnoreCase {
		cfg.IgnoreCase = false
	}

	fullLine, err := flags.GetBool("line-regexp")
	if err != nil {
		return err
	}

	if fullLine {
		cfg.WordsOnly = false
	}

	if ok := myio.ColorExists(cfg.Color); !ok {
		return ErrorUnknownColor
	}

	cfg.Pattern = adaptPattern(pattern, cfg)

	return nil
}

func adaptPattern(s string, cfg *Config) string {
	if cfg.FullLine {
		s = "^(" + s + ")$"
	}

	if cfg.WordsOnly {
		s = "\\b" + s + "\\b"
	}

	if cfg.IgnoreCase {
		s = "(?i)" + s
	}

	return s
}
