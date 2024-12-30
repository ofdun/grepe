package parsers

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

func init3DSlice(x, y, z int) [][][]int {
	slice := make([][][]int, x)
	for i := range slice {
		slice[i] = make([][]int, y)
		for j := range slice[i] {
			slice[i][j] = make([]int, z)
		}
	}

	return slice
}

func TestFullWord(t *testing.T) {
	regex := "test"
	text := []string{"test"}

	// Combine the []string into a single string (e.g., with spaces)
	combinedString := strings.Join(text, "\n")

	// Convert the string to an io.Reader
	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	matchingString, indexes, err := GetMatchIndexesArray(reader, re)

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

	// Combine the []string into a single string (e.g., with spaces)
	combinedString := strings.Join(text, "\n")

	// Convert the string to an io.Reader
	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	matchingString, indexes, err := GetMatchIndexesArray(reader, re)

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

	// Combine the []string into a single string (e.g., with spaces)
	combinedString := strings.Join(text, "\n")

	// Convert the string to an io.Reader
	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	matchingString, indexes, err := GetMatchIndexesArray(reader, re)

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

	// Combine the []string into a single string (e.g., with spaces)
	combinedString := strings.Join(text, "\n")

	// Convert the string to an io.Reader
	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	matchingString, indexes, err := GetMatchIndexesArray(reader, re)

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

	// Combine the []string into a single string (e.g., with spaces)
	combinedString := strings.Join(text, "\n")

	// Convert the string to an io.Reader
	reader := strings.NewReader(combinedString)

	re := regexp.MustCompile(regex)

	matchingString, indexes, err := GetMatchIndexesArray(reader, re)

	matchingStringOut := []string{"aboba"}
	indexesOut := init3DSlice(1, 1, 2)

	indexesOut[0][0][0] = 0
	indexesOut[0][0][1] = 5

	assert.ErrorIs(t, err, nil)
	assert.EqualValues(t, matchingString, &matchingStringOut)
	assert.EqualValues(t, indexes, &indexesOut)
}
