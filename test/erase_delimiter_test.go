package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEraseDelimiter(t *testing.T) {
	result := gotools.EraseDelimiter("Hello, world!", []string{","}, -1)

	assert.Equal(t, "Hello world!", result)
}
