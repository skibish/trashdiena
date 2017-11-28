package storage

import (
	"encoding/json"
	"fmt"

	"bitbucket.org/skibish/trashdiena/pkg/firebase"
)

const trashPath = "trash"

// Trash is struct shortuct
type Trash struct {
	firebase *firebase.Firebase
}

// TrashData decsibe trash data container
type TrashData struct {
	ID        string `json:"id"`
	Data      string `json:"data"`
	Published bool   `json:"published"`
}

// Set set trash to the database
func (t *Trash) Set(trash *TrashData) (err error) {
	refKey := fmt.Sprintf("%s/%s", trashPath, trash.ID)
	err = t.firebase.Set(refKey, trash)

	return
}

// GetNotPublished return all not published records
func (t *Trash) GetNotPublished() (finResult map[string]*TrashData, err error) {
	result, err := t.firebase.Get(trashPath)
	if err != nil {
		return
	}

	err = json.Unmarshal(result, &finResult)

	for k, v := range finResult {
		if v.Published {
			delete(finResult, k)
		}
	}

	return
}
