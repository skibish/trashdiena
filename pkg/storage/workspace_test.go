package storage

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestWorkspaceSet(t *testing.T) {
	fm := &mockFirebase{}
	var realRef string
	fm.set = func(path string, v interface{}) (err error) {
		realRef = path
		return nil
	}
	wp := &Workspace{firebase: fm}

	err := wp.Set(&WorkspaceData{ChannelID: "123", ID: "321", WebhookURL: "http://a/b/c"})
	if err != nil {
		t.Error(err)
	}
}

func TestWorkspaceSetErr(t *testing.T) {
	fm := &mockFirebase{}
	fm.set = func(path string, v interface{}) (err error) {
		return errors.New("Oops")
	}
	tr := &Workspace{firebase: fm}

	err := tr.Set(&WorkspaceData{ChannelID: "123", ID: "321", WebhookURL: "http://a/b/c"})
	if err.Error() != "Oops" {
		t.Error("Expected error, but everything looks good")
	}
}

func TestWorkspaceGetAll(t *testing.T) {
	fm := &mockFirebase{}
	fm.get = func(path string) (result json.RawMessage, err error) {
		return []byte(`{"12":{"id":"12","channel_id":"123","webhook_url":"http://a/b/c"},"23":{"id":"23","channel_id":"234","webhook_url":"http://d/e/f"}}`), nil
	}

	wp := &Workspace{firebase: fm}

	res, err := wp.GetAll()
	if err != nil {
		t.Error(err)
	}

	if len(res) != 2 {
		t.Error("Should be one element in the array")
	}
}

func TestWorkspaceGetAllErr(t *testing.T) {
	fm := &mockFirebase{}
	fm.get = func(path string) (result json.RawMessage, err error) {
		return []byte(``), errors.New("Oops")
	}

	wp := &Workspace{firebase: fm}

	_, err := wp.GetAll()
	if err.Error() != "Oops" {
		t.Error("Should be error, but everything is OK")
	}
}

func TestWorkspaceDelete(t *testing.T) {
	fm := &mockFirebase{}
	fm.delete = func(path string) (err error) {
		return nil
	}

	wp := Workspace{firebase: fm}
	err := wp.Delete(WorkspaceData{ChannelID: "123", ID: "321", WebhookURL: "http://a/b/c"})
	if err != nil {
		t.Errorf("Error was not expected: %v", err)
	}
}
