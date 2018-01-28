package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type callerHook struct{}

func (f *callerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (f *callerHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 1, 1)
	n := runtime.Callers(7, pc)
	if n == 0 {
		return nil
	}

	pc = pc[:n]
	frames := runtime.CallersFrames(pc)

	for {
		frame, more := frames.Next()

		if !strings.Contains(frame.File, "github.com/sirupsen/logrus") {
			entry.Data["file"] = fmt.Sprintf("%v:%v", frame.File, frame.Line)
			entry.Data["function"] = path.Base(frame.Function)
			break
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
