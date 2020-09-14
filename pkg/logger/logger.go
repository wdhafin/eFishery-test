package logger

import (
	"github.com/sirupsen/logrus"
)

// CaptureErr is
func CaptureErr(err error) {
	logrus.Error(err)
}
