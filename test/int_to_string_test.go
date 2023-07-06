package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToString(t *testing.T) {
	result := gotools.IntToString(666)

	assert.Equal(t, "666", result)
}
