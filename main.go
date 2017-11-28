package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/namsral/flag"

	"bitbucket.org/skibish/trashdiena/api"
	"bitbucket.org/skibish/trashdiena/firebase"
	"bitbucket.org/skibish/trashdiena/scheduler"
	"bitbucket.org/skibish/trashdiena/slack"
	"bitbucket.org/skibish/trashdiena/storage"
)

func main() {
	// API
	apiPort := flag.String("api-port", "80", "API port")

	// Slack
	clientID := flag.String("client-id", "", "Slack Client ID")
	clientSecret := flag.String("client-secret", "", "Slack Secret")
	redirectURL := flag.String("redirect-url", "", "specific redirect URL")

	// Firebase
	firebaseDB := flag.String("firebase-db", "", "Firebase DB URL")
	firebaseSecret := flag.String("firebase-secret", "firebase-secret.json", "Firebase Secret JSON")
	startAs := flag.String("start-as", "api", "Mode to start application as \"api\" or \"scheduler\"")

	flag.Parse()

	// Try to initialize all mandatory parts of the application
	fbase, err := firebase.New(*firebaseDB, *firebaseSecret)
	if err != nil {
		log.Fatal(err)
	}

	sc := slack.New(*clientID, *clientSecret, *redirectURL)
	sg := storage.New(fbase)

	var a *api.API
	// if we start application as API, then try to start it async
	if *startAs == "api" {
		a := api.New(sc, sg)
		go func() {
			log.Fatal(a.Start(*apiPort))
		}()
	} else {
		log.Println("Starting scheduler...")
		scheduler.New(sg).Start()
		log.Println("Scheduler started!")
	}

	// handle all the gracefull shutdowns
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	log.Println("Application started!")
	select {
	case <-sigs:
		log.Println("Performing shutdown...")
		if a != nil {
			a.Shutdown()
		}
		log.Println("Exiting...")
	}
}
