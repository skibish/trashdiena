package api

import (
	"io/ioutil"
	"log"
	"net/http"

	"bitbucket.org/skibish/trashdiena/storage"
	uuid "github.com/satori/go.uuid"
)

func (a *API) handlerInit(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	resp, err := a.slackClient.OAuthAccess(code)
	if err != nil {
		log.Printf("ERR: %v\n", err)
		return
	}

	err = a.db.Workspace.Set(&storage.WorkspaceData{
		ID:         resp.TeamID,
		WebhookURL: resp.WebhookURL,
		ChannelID:  resp.ChannelID,
	})

	if err != nil {
		log.Printf("ERR: %v\n", err)
	}

	http.Redirect(w, r, resp.RedirectURL, http.StatusTemporaryRedirect)
	return
}

// SHOULD NOT BE PUBLIC (currently)
func (a *API) handlerCreate(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ERR: %v\n", err)
		return
	}

	err = a.db.Trash.Set(&storage.TrashData{
		ID:        uuid.NewV4().String(),
		Data:      string(b),
		Published: false,
	})

	if err != nil {
		log.Printf("ERR: %v\n", err)
		return
	}

	return
}
