package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/skibish/trashdiena/pkg/config"

	"github.com/skibish/trashdiena/pkg/api"
	"github.com/skibish/trashdiena/pkg/firebase"
	"github.com/skibish/trashdiena/pkg/slack"
	"github.com/skibish/trashdiena/pkg/storage"
)

func main() {
	log.Println("Starting authorizer...")

	c := config.Parse()

	fbase, err := firebase.New(c.FirebaseDB, c.FirebaseSecretPath)
	if err != nil {
		log.Fatal(err)
	}

	sc := slack.New(c.ClientID, c.ClientSecret, c.RedirectURL)
	sg := storage.New(fbase)
	a := api.New(sc, sg)

	// shutdown gracefully
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
		<-sigs
		log.Println("Performing shutdown...")
		if err := a.Shutdown(); err != nil {
			log.Printf("Failed to shutdown server: %v", err)
		}
	}()

	log.Printf("Authorizer is ready to listen on port: %s", c.APIPort)
	if err := a.Start(c.APIPort); err != http.ErrServerClosed {
		log.Printf("Server failed: %v", err)
		os.Exit(1)
	}

	log.Println("Exiting...")
}
