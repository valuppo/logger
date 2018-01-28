package logger

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

type callerHook struct{}

func (f *callerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (f *callerHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	if n == 0 {
		return nil
	}

	pc = pc[:n]
	frames := runtime.CallersFrames(pc)

	for {
		frame, more := frames.Next()
		entry.Data["file"] = frame.File
		entry.Data["func"] = frame.Function
		entry.Data["line"] = frame.Line
		if !more {
			break
		}
	}
	return nil
}
