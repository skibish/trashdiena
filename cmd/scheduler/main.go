package main

import (
	"log"

	"bitbucket.org/skibish/trashdiena/pkg/config"
	"bitbucket.org/skibish/trashdiena/pkg/firebase"
	"bitbucket.org/skibish/trashdiena/pkg/scheduler"
	"bitbucket.org/skibish/trashdiena/pkg/storage"
)

func main() {
	log.Println("Starting scheduler...")
	c := config.Parse()

	// Try to initialize all mandatory parts of the application
	fbase, err := firebase.New(c.FirebaseDB, c.FirebaseSecretPath)
	if err != nil {
		log.Fatal(err)
	}

	scheduler.New(storage.New(fbase)).Start()
	log.Println("Scheduler started!")

	// TODO: Gracefull shutdown for the Scheduler
	select {}
}
