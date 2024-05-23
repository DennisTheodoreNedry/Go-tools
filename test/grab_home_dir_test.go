package test

import (
	gotools "github.com/DennisTheodoreNedry/Go-tools"
	"github.com/stretchr/testify/assert"
	"os/user"
	"testing"
)

func TestGrabHomeDir(t *testing.T) {
	result := gotools.GrabHomeDir()
	path, _ := user.Current()

	assert.Equal(t, path.HomeDir, result)
}
