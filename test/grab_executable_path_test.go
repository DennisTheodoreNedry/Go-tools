package test

import (
	gotools "github.com/DennisTheodoreNedry/Go-tools"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGrabExecutablePath(t *testing.T) {
	result := gotools.GrabExecutablePath()

	assert.Equal(t, os.Args[0], result)
}
