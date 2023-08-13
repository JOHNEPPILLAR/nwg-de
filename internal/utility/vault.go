// Package utility - Wrapper for Hashicorp vault
package utility

import (
	"errors"
	"os"

	"github.com/buger/jsonparser"
)

// GetVaultSecret - Get secret
func GetVaultSecret(secret string) (value string, err error) {

	filePath := "/app/config"
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	vaultValue, err := jsonparser.GetString(raw, "data", secret)
	if err != nil {
		return "", errors.New("No data found")
	}

	return vaultValue, nil
}

// UpdateVaultSecret - Update secret
func UpdateVaultSecret(secret string, value string) error {

	filePath := "/app/config"
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	newConfigFileData, err := jsonparser.Set(raw, []byte(value), "data", secret)
	if err != nil {
		return errors.New("Unable to update vault")
	}

	err = os.WriteFile(filePath, []byte(newConfigFileData), 0)
	if err != nil {
		return errors.New("Unable to update vault")
	}

	return nil
}
