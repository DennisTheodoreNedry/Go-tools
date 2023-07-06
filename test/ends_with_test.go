package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndsWith(t *testing.T) {
	result := gotools.EndsWith("Hello, world!", []string{"!"})

	ok := result["!"]

	assert.Equal(t, true, ok)
}

func TestDoesntEndWith(t *testing.T) {
	result := gotools.EndsWith("Hello, world!", []string{"H"})

	ok := result["H"]

	assert.Equal(t, false, ok)
}
