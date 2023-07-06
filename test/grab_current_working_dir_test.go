package test

import (
	gotools "github.com/s9rA16Bf4/Go-tools"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGrabCWD(t *testing.T) {
	result := gotools.GrabCWD()
	path, _ := os.Getwd()

	assert.Equal(t, path, result)
}
