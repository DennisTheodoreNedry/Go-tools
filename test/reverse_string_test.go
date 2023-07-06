package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	value := "mephisto"
	gotools.ReverseString(&value)

	assert.Equal(t, "otsihpem", value)
}
