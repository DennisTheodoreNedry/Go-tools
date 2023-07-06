package test

import (
	gotools "github.com/s9rA16Bf4/Go-tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToBool(t *testing.T) {
	result := gotools.StringToBoolean("true")

	assert.Equal(t, true, result)

	result = gotools.StringToBoolean("false")

	assert.Equal(t, false, result)
}
