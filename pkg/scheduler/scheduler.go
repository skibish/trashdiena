package scheduler

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/skibish/trashdiena/pkg/storage"
)

// Scheduler describe scheduler
type Scheduler struct {
	sg             *storage.Storage
	publishDay     time.Weekday
	maxToPublish   int
	countPublished int
}

// New return new Scheduler
func New(storage *storage.Storage) *Scheduler {
	return &Scheduler{
		sg:           storage,
		publishDay:   time.Wednesday,
		maxToPublish: 1,
	}
}

// Start starts server
func (s *Scheduler) Start() {
	tick := time.NewTicker(time.Hour).C

	go func() {
		for {
			select {
			case t := <-tick:

				// If today is not the publish Day, reset all the information
				if s.publishDay != t.Weekday() {
					s.countPublished = 0
				}

				if s.allow(t) {
					log.Printf("It's %s! Publishing some trash...", s.publishDay.String())
					if err := s.publish(); err != nil {
						log.Println(err)
					}
					s.countPublished++
				}
			}
		}

	}()
}

func (s *Scheduler) allow(t time.Time) bool {
	if !(t.Hour() > 10 && t.Hour() <= 22) {
		return false
	}

	if s.publishDay != t.Weekday() {
		return false
	}

	if s.countPublished >= s.maxToPublish {
		return false
	}

	return true
}

// Publish publish some trash to all the channels
func (s *Scheduler) publish() (err error) {
	workspaces, err := s.sg.Workspace.GetAll()
	if err != nil {
		return
	}

	trashs, err := s.sg.Trash.GetNotPublished()
	if err != nil {
		return
	}

	// get not published trash, mark it as published and
	// save to the database
	var trash string
	for _, v := range trashs {
		trash = v.Data
		v.Published = true
		s.sg.Trash.Set(v)
		break
	}

	// send trash to all subscribed workspaces
	for _, v := range workspaces {
		go func(wd storage.WorkspaceData) {
			json := []byte(fmt.Sprintf("{\"text\":%q}", trash))
			_, err := http.Post(wd.WebhookURL, "application/json", bytes.NewBuffer(json))
			if err != nil {
				log.Printf("Deleting workspace \"%s:%s\" because failed to POST", wd.ChannelID, wd.ID)
				s.sg.Workspace.Delete(wd)
				return
			}
		}(v)
	}

	return
}
