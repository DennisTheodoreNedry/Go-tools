package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToBool(t *testing.T) {
	result := gotools.StringToBool("true")

	assert.Equal(t, true, result)

	result = gotools.StringToBool("false")

	assert.Equal(t, false, result)
}
