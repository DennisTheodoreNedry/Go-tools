package test

import (
	gotools "github.com/s9rA16Bf4/Go-tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitString(t *testing.T) {
	value := gotools.SplitString("Hello")

	assert.Equal(t, []string{"H", "e", "l", "l", "o"}, value)
}
