package api

import (
	"context"
	"net/http"

	"bitbucket.org/skibish/trashdiena/pkg/slack"
	"bitbucket.org/skibish/trashdiena/pkg/storage"
)

// API is a structure that contains API for the bot
type API struct {
	slackClient *slack.Slack
	db          *storage.Storage
	server      *http.Server
}

// New return new API instance
func New(slackClient *slack.Slack, db *storage.Storage) *API {
	return &API{
		slackClient: slackClient,
		db:          db,
	}
}

// Start starts the API server
func (a *API) Start(port string) error {
	s := &http.Server{
		Addr:    ":" + port,
		Handler: a.bootRouter(),
	}

	a.server = s

	return s.ListenAndServe()
}

// Shutdown performs graceful API shutdown
func (a *API) Shutdown() error {
	return a.server.Shutdown(context.Background())
}
