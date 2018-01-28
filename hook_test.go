package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func BenchmarkFire(b *testing.B) {
	logrus.AddHook(NewCallerHook())
	logrus.SetLevel(logrus.ErrorLevel)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logrus.Info("test")
	}
}
