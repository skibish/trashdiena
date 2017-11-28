package main

import (
	"log"
	"os"
	"os/signal"

	"bitbucket.org/skibish/trashdiena/pkg/config"

	"bitbucket.org/skibish/trashdiena/pkg/api"
	"bitbucket.org/skibish/trashdiena/pkg/firebase"
	"bitbucket.org/skibish/trashdiena/pkg/slack"
	"bitbucket.org/skibish/trashdiena/pkg/storage"
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
