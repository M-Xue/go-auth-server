package server

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

// Guide for zerolog: https://betterstack.com/community/guides/logging/zerolog/
func initLogger(environment string) (zerolog.Logger, error) {
	var logger zerolog.Logger
	switch environment {
	case "development":
		file, err := os.OpenFile(
			"server.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		if err != nil {
			return logger, err
		}

		logger = zerolog.New(file).
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
