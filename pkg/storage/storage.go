package storage

import "bitbucket.org/skibish/trashdiena/pkg/firebase"

// Storage describe storage
type Storage struct {
	Workspace *Workspace
	Trash     *Trash
}

// New return new Storage
func New(firebase firebase.IFirebase) *Storage {
	return &Storage{
		Workspace: &Workspace{firebase: firebase},
		Trash:     &Trash{firebase: firebase},
	}
}
