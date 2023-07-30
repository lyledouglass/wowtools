package test

import (
	"testing"
	"wowtools/pkg/utilities"
)

func TestLogrus(t *testing.T) {
	t.Skip()
	utilities.SetupLogger("info")
	utilities.Log.Info("Testing Logrus")
}
