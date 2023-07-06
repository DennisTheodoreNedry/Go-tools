package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartsWith(t *testing.T) {
	result := gotools.StartsWith("Hello my name is XXXAAAXXX", []string{"Hello"})

	ok := result["Hello"]

	assert.Equal(t, true, ok)
}

func TestDoesntStartWith(t *testing.T) {
	result := gotools.StartsWith("Hello, world!", []string{"Peter"})

	ok := result["Peter"]

	assert.Equal(t, false, ok)
}
