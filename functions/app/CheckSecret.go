package app

import "strings"

// CheckSecret verifies if provided secret matches provided secret keys
func (app *App) CheckSecret(secretKey string) bool {
	allowedSecretKeys := strings.Split(app.Config.SecretKeys, ",")
	for _, allowedKey := range allowedSecretKeys {
		if allowedKey == secretKey {
			return true
		}
	}
	return false
}
