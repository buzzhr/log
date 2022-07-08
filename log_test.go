package log

import (
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	Info("this is info log.")
	PrintError(errors.New("xxx"))
	Infof("this is infof %s.", "log")
	// Fatal("fatalxxxx.")
}
