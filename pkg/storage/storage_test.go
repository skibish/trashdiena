package storage

import (
	"encoding/json"
	"testing"
)

type mockFirebase struct {
	set         func(path string, v interface{}) (err error)
	get         func(path string) (result json.RawMessage, err error)
	filterEqual func(path, field string, value interface{}) (result json.RawMessage, err error)
}

func (m mockFirebase) Set(path string, v interface{}) (err error) {
	return m.set(path, v)
}
func (m mockFirebase) Get(path string) (result json.RawMessage, err error) {
	return m.get(path)
}

func (m mockFirebase) FilterEqual(path, field string, value interface{}) (result json.RawMessage, err error) {
	return m.filterEqual(path, field, value)
}

func TestNew(t *testing.T) {
	fm := &mockFirebase{}

	// check that function works without any compile errors, etc.
	New(fm)
}
