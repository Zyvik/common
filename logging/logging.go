package logging

import (
	"bufio"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// logging literals
const (
	TraceID = "traceID"
	SpanID  = "spanID"
)

type Config struct {
	Level      string `env:"LEVEL,required"`
	OutputFile string `env:"OUTPUT_FILE"`
}

func InitLogger(cfg *Config) *logrus.Logger {
	log := logrus.New()
	logLevel := logrus.DebugLevel // default value
	if parsedLevel, err := logrus.ParseLevel(cfg.Level); err == nil {
		logLevel = parsedLevel
	}
	log.SetLevel(logLevel)
	log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano})
	if cfg.OutputFile != "" {
		file, err := os.OpenFile(cfg.OutputFile, os.O_CREATE, os.ModeAppend)
		if err != nil {
			log.WithError(err).Error("Failed to set logging output file")
			return log
		}
		log.SetOutput(bufio.NewWriter(file))
	}
	return log
}
