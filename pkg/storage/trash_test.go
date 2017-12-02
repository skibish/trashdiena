package storage

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestTrashSet(t *testing.T) {
	fm := &mockFirebase{}
	var realRef string
	fm.set = func(path string, v interface{}) (err error) {
		realRef = path
		return nil
	}
	tr := &Trash{firebase: fm}

	err := tr.Set(&TrashData{Data: "ooomg", ID: "123", Published: false})
	if err != nil {
		t.Error(err)
	}
}

func TestTrashSetErr(t *testing.T) {
	fm := &mockFirebase{}
	fm.set = func(path string, v interface{}) (err error) {
		return errors.New("Oops")
	}
	tr := &Trash{firebase: fm}

	err := tr.Set(&TrashData{Data: "ooomg", ID: "123", Published: false})
	if err.Error() != "Oops" {
		t.Error("Expected error, but everything looks good")
	}
}

func TestGetNotPublished(t *testing.T) {
	fm := &mockFirebase{}
	fm.get = func(path string) (result json.RawMessage, err error) {
		return []byte(`{"12":{"id":"12","data":"aaa","published":true},"23":{"id":"23","data":"bbb","published":false}}`), nil
	}

	tr := &Trash{firebase: fm}

	res, err := tr.GetNotPublished()
	if err != nil {
		t.Error(err)
	}

	if len(res) != 1 {
		t.Error("Should be one element in the array")
	}

}

func TestGetNotPublishedErr(t *testing.T) {
	fm := &mockFirebase{}
	fm.get = func(path string) (result json.RawMessage, err error) {
		return []byte(``), errors.New("Oops")
	}

	tr := &Trash{firebase: fm}

	_, err := tr.GetNotPublished()
	if err.Error() != "Oops" {
		t.Error("Expected error, but everything is OK")
	}

}
