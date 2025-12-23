package config

import (
	"errors"
	"os"
)

// LoadCredentials reads LinkedIn credentials from environment variables
func LoadCredentials() (string, string, error) {
	email := os.Getenv("LINKEDIN_EMAIL")
	password := os.Getenv("LINKEDIN_PASSWORD")

	if email == "" || password == "" {
		return "", "", errors.New("LINKEDIN_EMAIL or LINKEDIN_PASSWORD not set")
	}

	return email, password, nil
}
