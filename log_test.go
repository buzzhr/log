package defaultlog

import (
	"testing"
)

func TestRun(t *testing.T) {
	Info("this is info log.")
	Infof("this is infof %s.", "log")
	Fatal("fatal.")
}
