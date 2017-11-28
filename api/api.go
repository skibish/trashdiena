package api

import (
	"context"
	"log"
	"net/http"

	"bitbucket.org/skibish/trashdiena/slack"
	"bitbucket.org/skibish/trashdiena/storage"
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

// bootRouter boots router
func (a *API) bootRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.router)

	return mux
}

// router for the requests
func (a *API) router(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, " ", r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/init":
			a.handlerInit(w, r)
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message":"not found"}`))
		}
	}
}
