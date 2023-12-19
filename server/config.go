package server

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUrl               string
	TokenSymmetricKey   string
	AccessTokenDuration time.Duration
	ClientUrl           string
}

func envError(envName string) error {
	return fmt.Errorf("environment variable missing: %s", envName)
}

func loadConfig() (*Config, error) {
	godotenv.Load(".env")

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return nil, envError("DB_URL")
	}

	tokenSymmetricKey := os.Getenv("TOKEN_SYMMETRIC_KEY")
	if tokenSymmetricKey == "" {
		return nil, envError("TOKEN_SYMMETRIC_KEY")
	}

	accessTokenDurationString := os.Getenv("ACCESS_TOKEN_DURATION")
	if accessTokenDurationString == "" {
		return nil, envError("ACCESS_TOKEN_DURATION")
	}

	accessTokenDuration, err := time.ParseDuration(accessTokenDurationString)
	if err != nil {
		return nil, fmt.Errorf(
			"environment variable (%s) parsing error: %w",
			"ACCESS_TOKEN_DURATION",
			err,
		)
	}

	clientUrl := os.Getenv("CLIENT_URL")
	if clientUrl == "" {
		return nil, envError("CLIENT_URL")
	}

	return &Config{
		DbUrl:               dbUrl,
		TokenSymmetricKey:   tokenSymmetricKey,
		AccessTokenDuration: accessTokenDuration,
		ClientUrl:           clientUrl,
	}, nil
}
