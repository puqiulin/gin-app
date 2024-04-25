package logger

import (
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type FunctionNameHook struct{}

func (hook *FunctionNameHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *FunctionNameHook) Fire(entry *logrus.Entry) error {
	pc, _, _, ok := runtime.Caller(8) // Adjust the call depth as needed
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			entry.Data["function"] = strings.TrimPrefix(fn.Name(), "main.")
		}
	}
	return nil
}
