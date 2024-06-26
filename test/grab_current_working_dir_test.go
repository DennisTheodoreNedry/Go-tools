package test

import (
	gotools "github.com/DennisTheodoreNedry/Go-tools"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGrabCWD(t *testing.T) {
	result := gotools.GrabCWD()
	path, _ := os.Getwd()

	assert.Equal(t, path, result)
}
