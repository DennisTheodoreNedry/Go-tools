package gotools

import (
	"os/user"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Grabs the current users home directory
func GrabHomeDir() string {
	path, err := user.Current()

	if err != nil {
		notify.Error(err.Error(), "GrabHomeDir()")
	}

	return path.HomeDir
}
