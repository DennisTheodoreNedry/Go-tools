package test

import (
	gotools "github.com/s9rA16Bf4/Go-tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEraseDelimiter(t *testing.T) {
	result := gotools.EraseDelimiter("Hello, world!", []string{","}, -1)

	assert.Equal(t, "Hello world!", result)
}
