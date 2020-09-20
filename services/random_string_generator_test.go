package services

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomStringGenerator(t *testing.T) {
	result := RandomStringGenerator(5)
	assert.Equal(t, len(result), 5)

	stringLetters := string(letters)
	for i := range result {
		assert.True(t, strings.ContainsRune(stringLetters, rune(result[i])))
	}
}
