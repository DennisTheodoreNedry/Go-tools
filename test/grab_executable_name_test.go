package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrabExecutableName(t *testing.T) {
	result := gotools.GrabExecutableName()

	assert.Equal(t, filepath.Base(os.Args[0]), result)
}
