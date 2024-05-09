package log

import (
	"os"

	"github.com/google/wire"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var ProviderSet = wire.NewSet(NewLog)

var Log *logrus.Logger

func NewLog(w *kafka.Writer) *logrus.Logger {
	Log = logrus.New()

	Log.AddHook(&FunctionNameHook{})
	//Log.AddHook(&KafkaHook{Writer: w})

	Log.Out = os.Stdout
	Log.Formatter = &logrus.JSONFormatter{}
	Log.Level = logrus.DebugLevel

	return Log
}
