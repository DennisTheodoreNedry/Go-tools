package test

import (
	"testing"

	gotools "github.com/s9rA16Bf4/Go-tools"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	result := gotools.Contains("Hello, world!", []string{","})

	ok := result[","]

	assert.Equal(t, true, ok)
}

func TestDoestNotContainValue(t *testing.T) {
	result := gotools.Contains("Hello, world!", []string{"."})

	ok := result[","]

	assert.Equal(t, false, ok)
}
