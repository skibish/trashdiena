package config

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {

	envs := map[string]string{
		"API_PORT":        "99",
		"CLIENT_ID":       "abc",
		"CLIENT_SECRET":   "secret",
		"REDIRECT_URL":    "http://redirect",
		"FIREBASE_DB":     "http://fbdb",
		"FIREBASE_SECRET": "fsecret",
		"TRASH_FILE":      "trash.scv",
	}

	for k, v := range envs {
		os.Setenv(k, v)
	}

	c := Parse()

	if c.APIPort != envs["API_PORT"] {
		t.Errorf("Expected APIPort=%s, got %s", envs["API_PORT"], c.APIPort)
	}

	if c.FirebaseSecretPath != envs["FIREBASE_SECRET"] {
		t.Errorf("Expected FirebaseSecretPath=%s, got %s", envs["FIREBASE_SECRET"], c.FirebaseSecretPath)
	}
}
