package test

import (
	gotools "github.com/s9rA16Bf4/Go-tools"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestGrabExecutableName(t *testing.T) {
	result := gotools.GrabExecutableName()

	assert.Equal(t, filepath.Base(os.Args[0]), result)
}
