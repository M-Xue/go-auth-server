package server

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Guide for zerolog: https://betterstack.com/community/guides/logging/zerolog/
func initLogger(environment string) (zerolog.Logger, error) {
	// TODO set up loggin with sqlc. im pre sure there is a way to use sqlc with zerolog

	var logger zerolog.Logger
	switch environment {
	case "development":
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
			FormatMessage: func(i interface{}) string {
				return fmt.Sprintf("| %s |", i)
			},
		}

		logger = zerolog.New(consoleWriter).
			With().
			Timestamp().
			Logger()

	case "production":
		logger = zerolog.New(os.Stdout).
			With().
			Timestamp().
			Logger()

	default:
		return zerolog.Logger{}, fmt.Errorf("invalid environment variables")
	}

	return logger, nil
}
