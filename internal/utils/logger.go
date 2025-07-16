package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// NewLogger initializes and returns a zerolog.Logger instance
func NewLogger() zerolog.Logger {
	// Set global time format
	zerolog.TimeFieldFormat = time.RFC3339

	// Set log level based on env variable
	logLevel := zerolog.InfoLevel
	if os.Getenv("DEBUG") == "1" {
		logLevel = zerolog.DebugLevel
	}

	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Caller(). 
		Logger().
		Level(logLevel)

	log.Logger = logger // Replace global logger 

	return logger
}
