package utilities

import (
	"strings"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func SetupLogger(loglevel string) {
	Log.SetLevel(getLogLevelFromString(loglevel))
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    loglevel == "trace",
		DisableTimestamp: loglevel != "trace",
	})
}
func getLogLevelFromString(loglevel string) logrus.Level {
	switch strings.ToLower(loglevel) {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}
