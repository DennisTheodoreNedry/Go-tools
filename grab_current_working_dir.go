package gotools

import (
	"os"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Grabs the current working path
func GrabCWD() string {
	path, err := os.Getwd()
	if err != nil {
		notify.Error(err.Error(), "GrabCWD()")
	}
	return path
}
