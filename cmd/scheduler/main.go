package main

import (
	"log"

	"github.com/skibish/trashdiena/pkg/config"
	"github.com/skibish/trashdiena/pkg/firebase"
	"github.com/skibish/trashdiena/pkg/scheduler"
	"github.com/skibish/trashdiena/pkg/storage"
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
