package parsers

import (
	"github.com/stretchr/testify/assert"
	"grepe/internal/config"
	"regexp"
	"strings"
	"testing"
)

var (
	defaultConfig = &config.Config{
		MaxLines: -1,
	}
)

func init3DSlice(x, y, _z int) [][][]int {
	slice := make([][][]int, x)
	for i := range slice {
		slice[i] = make([][]int, y)
		for j := range slice[i] {
			slice[i][j] = make([]int, _z)
		}
	}

	return slice
}

func TestFullWord(t *testing.T) {
	regex := "test"
	text := []string{"test"}

	combinedString := strings.Join(text, "\n")

	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	matchingString, indexes, err := GetMatchIndexesArray(reader, re, defaultConfig)

	matchingStringOut := text
	indexesOut := init3DSlice(1, 1, 2)

	indexesOut[0][0][0] = 0
	indexesOut[0][0][1] = 4

	assert.ErrorIs(t, err, nil)
	assert.EqualValues(t, matchingString, &matchingStringOut)
	assert.EqualValues(t, indexes, &indexesOut)
}

func TestStartOfWord(t *testing.T) {
	regex := "w"
	text := []string{"word"}

	combinedString := strings.Join(text, "\n")

	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	matchingString, indexes, err := GetMatchIndexesArray(reader, re, defaultConfig)

	matchingStringOut := text
	indexesOut := init3DSlice(1, 1, 2)

	indexesOut[0][0][0] = 0
	indexesOut[0][0][1] = 1

	assert.ErrorIs(t, err, nil)
	assert.EqualValues(t, matchingString, &matchingStringOut)
	assert.EqualValues(t, indexes, &indexesOut)
}

func TestEndOfWord(t *testing.T) {
	regex := "d"
	text := []string{"word"}

	combinedString := strings.Join(text, "\n")

	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	matchingString, indexes, err := GetMatchIndexesArray(reader, re, defaultConfig)

	matchingStringOut := text
	indexesOut := init3DSlice(1, 1, 2)

	indexesOut[0][0][0] = 3
	indexesOut[0][0][1] = 4

	assert.ErrorIs(t, err, nil)
	assert.EqualValues(t, matchingString, &matchingStringOut)
	assert.EqualValues(t, indexes, &indexesOut)
}

func TestFewWords(t *testing.T) {
	regex := "ab"
	text := []string{"word", "aboba", "bebra"}

	combinedString := strings.Join(text, "\n")

	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	matchingString, indexes, err := GetMatchIndexesArray(reader, re, defaultConfig)

	matchingStringOut := []string{"aboba"}
	indexesOut := init3DSlice(1, 1, 2)

	indexesOut[0][0][0] = 0
	indexesOut[0][0][1] = 2

	assert.ErrorIs(t, err, nil)
	assert.EqualValues(t, matchingString, &matchingStringOut)
	assert.EqualValues(t, indexes, &indexesOut)
}

func TestMoreComplexRegex(t *testing.T) {
	regex := "^a[a-z]+a$"
	text := []string{"word", "aboba", "bebra"}

	combinedString := strings.Join(text, "\n")

	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	matchingString, indexes, err := GetMatchIndexesArray(reader, re, defaultConfig)

	matchingStringOut := []string{"aboba"}
	indexesOut := init3DSlice(1, 1, 2)

	indexesOut[0][0][0] = 0
	indexesOut[0][0][1] = 5

	assert.ErrorIs(t, err, nil)
	assert.EqualValues(t, matchingString, &matchingStringOut)
	assert.EqualValues(t, indexes, &indexesOut)
}

func TestFewWordsMatched(t *testing.T) {
	regex := "^a[a-z]+a$"
	text := []string{"word", "aboba", "abaunda"}

	combinedString := strings.Join(text, "\n")

	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	matchingString, indexes, err := GetMatchIndexesArray(reader, re, defaultConfig)

	matchingStringOut := []string{"aboba", "abaunda"}
	indexesOut := init3DSlice(2, 1, 2)

	indexesOut[0][0][0] = 0
	indexesOut[0][0][1] = 5

	indexesOut[1][0][0] = 0
	indexesOut[1][0][1] = 7

	assert.ErrorIs(t, err, nil)
	assert.EqualValues(t, matchingString, &matchingStringOut)
	assert.EqualValues(t, indexes, &indexesOut)
}

func TestFewMatchesInWord(t *testing.T) {
	regex := "(ab|ba)"
	text := []string{"word", "aboba"}

	combinedString := strings.Join(text, "\n")

	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	matchingString, indexes, err := GetMatchIndexesArray(reader, re, defaultConfig)

	matchingStringOut := []string{"aboba"}
	indexesOut := init3DSlice(1, 2, 2)

	indexesOut[0][0][0] = 0
	indexesOut[0][0][1] = 2

	indexesOut[0][1][0] = 3
	indexesOut[0][1][1] = 5

	assert.ErrorIs(t, err, nil)
	assert.EqualValues(t, matchingString, &matchingStringOut)
	assert.EqualValues(t, indexes, &indexesOut)
}

// Using inverted flag
func TestInverted(t *testing.T) {
	regex := "(ab|ba)"
	text := []string{"word", "aboba"}

	combinedString := strings.Join(text, "\n")

	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	defaultConfig.Inverted = true
	matchingString, indexes, err := GetMatchIndexesArray(reader, re, defaultConfig)
	defaultConfig.Inverted = false

	matchingStringOut := []string{"word", "aboba"}
	indexesOut := init3DSlice(2, 1, 2)

	indexesOut[0][0][0] = 0
	indexesOut[0][0][1] = 4

	indexesOut[1][0][0] = 2
	indexesOut[1][0][1] = 3

	assert.ErrorIs(t, err, nil)
	assert.EqualValues(t, &matchingStringOut, matchingString)
	assert.EqualValues(t, &indexesOut, indexes)
}

func TestInvertedNotEndOfLine(t *testing.T) {
	regex := "(ab|ba)"
	text := []string{"word", "aboba", "abobaa"}

	combinedString := strings.Join(text, "\n")

	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	defaultConfig.Inverted = true
	matchingString, indexes, err := GetMatchIndexesArray(reader, re, defaultConfig)
	defaultConfig.Inverted = false

	matchingStringOut := []string{"word", "aboba", "abobaa"}
	indexesOut := init3DSlice(3, 1, 2)

	indexesOut[0][0][0] = 0
	indexesOut[0][0][1] = 4

	indexesOut[1][0][0] = 2
	indexesOut[1][0][1] = 3

	indexesOut[2][0][0] = 2
	indexesOut[2][0][1] = 3
	indexesOut[2] = append(indexesOut[2], make([]int, 1))
	indexesOut[2][1] = make([]int, 2)
	indexesOut[2][1][0] = 5
	indexesOut[2][1][1] = 6

	assert.ErrorIs(t, err, nil)
	assert.EqualValues(t, &matchingStringOut, matchingString)
	assert.EqualValues(t, &indexesOut, indexes)
}

func TestMaxLinesValid(t *testing.T) {
	regex := "ab"
	text := []string{"absolute", "aboba"}

	combinedString := strings.Join(text, "\n")

	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	defaultConfig.MaxLines = 1
	matchingString, indexes, err := GetMatchIndexesArray(reader, re, defaultConfig)
	defaultConfig.MaxLines = -1

	matchingStringOut := []string{"absolute"}
	indexesOut := init3DSlice(1, 1, 2)

	indexesOut[0][0][0] = 0
	indexesOut[0][0][1] = 2

	assert.ErrorIs(t, err, nil)
	assert.EqualValues(t, matchingString, &matchingStringOut)
	assert.EqualValues(t, indexes, &indexesOut)
}

func TestMaxLinesInvalid(t *testing.T) {
	regex := "ab"
	text := []string{"bebra", "absolute", "aboba"}

	combinedString := strings.Join(text, "\n")

	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	defaultConfig.MaxLines = 2
	matchingString, indexes, err := GetMatchIndexesArray(reader, re, defaultConfig)
	defaultConfig.MaxLines = -1

	matchingStringOut := []string{"absolute"}
	indexesOut := init3DSlice(1, 1, 2)

	indexesOut[0][0][0] = 0
	indexesOut[0][0][1] = 2

	assert.ErrorIs(t, err, nil)
	assert.EqualValues(t, matchingString, &matchingStringOut)
	assert.EqualValues(t, indexes, &indexesOut)
}
