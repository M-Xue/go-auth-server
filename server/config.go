package server

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUrl             string
	ClientUrl         string
	TokenSymmetricKey string
	AccessDuration    string
	Environment       string
}

func envError(envName string) error {
	return fmt.Errorf("environment variable missing: %s", envName)
}

func initConfig() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, fmt.Errorf("loading .env failed")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return Config{}, envError("DB_URL")
	}

	clientUrl := os.Getenv("CLIENT_URL")
	if clientUrl == "" {
		return Config{}, envError("CLIENT_URL")
	}

	tokenSymmetricKey := os.Getenv("TOKEN_SYMMETRIC_KEY")
	if tokenSymmetricKey == "" {
		return Config{}, envError("TOKEN_SYMMETRIC_KEY")
	}

	accessDuration := os.Getenv("ACCESS_TOKEN_DURATION")
	if accessDuration == "" {
		return Config{}, envError("ACCESS_TOKEN_DURATION")
	}
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		return Config{}, envError("ENVIRONMENT")
	}
	if environment != "production" && environment != "development" {
		return Config{}, fmt.Errorf("environment value invlid")
	}

	return Config{
		ClientUrl:         clientUrl,
		DbUrl:             dbUrl,
		TokenSymmetricKey: tokenSymmetricKey,
		AccessDuration:    accessDuration,
		Environment:       environment,
	}, nil
}
