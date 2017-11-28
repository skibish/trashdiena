package storage

import (
	"encoding/json"
	"fmt"

	"bitbucket.org/skibish/trashdiena/pkg/firebase"
)

const workspacePath = "workspace"

// Workspace describes all the configuration of the workspace
type Workspace struct {
	firebase *firebase.Firebase
}

// WorkspaceData decsibe WorkspaceData info
type WorkspaceData struct {
	ID         string `json:"id"`
	ChannelID  string `json:"channel_id"`
	WebhookURL string `json:"webhook_url"`
}

// Set creates or updates the workspace data
func (w *Workspace) Set(wd *WorkspaceData) (err error) {
	refKey := fmt.Sprintf("%s/%s:%s", workspacePath, wd.ID, wd.ChannelID)
	err = w.firebase.Set(refKey, wd)

	return
}

// GetAll return all workspaces
func (w *Workspace) GetAll() (finResult map[string]WorkspaceData, err error) {
	result, err := w.firebase.Get(workspacePath)
	if err != nil {
		return
	}

	err = json.Unmarshal(result, &finResult)

	return
}
