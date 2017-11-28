package config

import (
	"github.com/namsral/flag"
)

// Configuration describe all available configuration options
type Configuration struct {
	APIPort             string
	ClientID            string
	ClientSecret        string
	RedirectURL         string
	FirebaseDB          string
	FirebaseSecretPath  string
	PathToFileWithTrash string
}

// TODO: Each cmd can have it's own flags, currently we parse all for everyone wich is not user friendly

// Parse parses all the flags and environment variables
func Parse() (c Configuration) {
	// API
	flag.StringVar(&c.APIPort, "api-port", "80", "API port")

	// Slack
	flag.StringVar(&c.ClientID, "client-id", "", "Slack Client ID")
	flag.StringVar(&c.ClientSecret, "client-secret", "", "Slack Secret")
	flag.StringVar(&c.RedirectURL, "redirect-url", "", "specific redirect URL")

	// Firebase
	flag.StringVar(&c.FirebaseDB, "firebase-db", "", "Firebase DB URL")
	flag.StringVar(&c.FirebaseSecretPath, "firebase-secret", "firebase-secret.json", "Firebase Secret JSON")

	// File with trash
	flag.StringVar(&c.PathToFileWithTrash, "trash-file", "trash.csv", "File with trash. Each line in file represent some trash.")

	flag.Parse()

	return
}
