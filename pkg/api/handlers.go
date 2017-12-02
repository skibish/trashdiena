package api

import (
	"log"
	"net/http"

	"github.com/skibish/trashdiena/pkg/storage"
)

func (a *API) handlerInit(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		return
	}

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
