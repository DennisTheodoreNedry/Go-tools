package test

import (
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrabUsername(t *testing.T) {
	result := gotools.GrabUsername()
	user, _ := user.Current()

	assert.Equal(t, user.Username, result)
}
