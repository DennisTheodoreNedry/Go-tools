package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitString(t *testing.T) {
	value := gotools.SplitString("Hello")

	assert.Equal(t, []string{"H", "e", "l", "l", "o"}, value)
}
