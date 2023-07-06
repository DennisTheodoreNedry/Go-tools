package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrabCWD(t *testing.T) {
	result := gotools.GrabCWD()
	path, _ := os.Getwd()

	assert.Equal(t, path, result)
}
