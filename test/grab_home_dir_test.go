package test

import (
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrabHomeDir(t *testing.T) {
	result := gotools.GrabHomeDir()
	path, _ := user.Current()

	assert.Equal(t, path.HomeDir, result)
}
