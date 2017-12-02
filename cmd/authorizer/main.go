package main

import (
	"log"
	"os"
	"os/signal"

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
	go func() {
		log.Fatal(a.Start(c.APIPort))
	}()

	// handle all the gracefull shutdowns
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	select {
	case <-sigs:
		log.Println("Performing shutdown...")
		if a != nil {
			a.Shutdown()
		}
		log.Println("Exiting...")
	}
}
