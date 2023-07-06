package test

import (
	gotools "github.com/s9rA16Bf4/Go-tools"
	"github.com/stretchr/testify/assert"
	"testing"
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
