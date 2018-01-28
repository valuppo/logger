package logger

import (
	"fmt"
	"runtime"
	"strings"

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

		if !strings.Contains(frame.File, "runtime") &&
			!strings.Contains(frame.File, "github.com/valuppo/logger") &&
			!strings.Contains(frame.File, "github.com/sirupsen/logrus") {
			entry.Data["file"] = fmt.Sprintf("%v:%v", frame.File, frame.Line)
			entry.Data["function"] = frame.Function
		}

		if !more {
			break
		}
	}
	return nil
}

func NewCallerHook() *callerHook {
	return &callerHook{}
}
