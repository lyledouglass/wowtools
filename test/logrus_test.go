package test

import (
	"testing"
	"wowtools/internal/utilities"
)

func TestLogrus(t *testing.T) {
	t.Skip()
	utilities.SetupLogger("info")
	utilities.Log.Info("Testing Logrus")
}
