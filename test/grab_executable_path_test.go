package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrabExecutablePath(t *testing.T) {
	result := gotools.GrabExecutablePath()

	assert.Equal(t, os.Args[0], result)
}
