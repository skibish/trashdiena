package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"bitbucket.org/skibish/trashdiena/pkg/slack"
	"bitbucket.org/skibish/trashdiena/pkg/storage"
)

type mockFirebase struct {
	set func(path string, v interface{}) (err error)
	get func(path string) (result json.RawMessage, err error)
}

func (m mockFirebase) Set(path string, v interface{}) (err error) {
	return m.set(path, v)
}
func (m mockFirebase) Get(path string) (result json.RawMessage, err error) {
	return m.get(path)
}

type mockSlack struct {
	oAuthAccess func(code string) (r *slack.OAuthAccessResponse, err error)
}

func (m mockSlack) OAuthAccess(code string) (r *slack.OAuthAccessResponse, err error) {
	return m.oAuthAccess(code)
}

type redirectHandler struct{}

func (rh redirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "redirected")
}

func TestInitHandler(t *testing.T) {
	fm := mockFirebase{}
	fm.set = func(path string, v interface{}) (err error) {
		return nil
	}

	redirectServer := httptest.NewServer(&redirectHandler{})
	defer redirectServer.Close()

	sm := mockSlack{}
	sm.oAuthAccess = func(code string) (r *slack.OAuthAccessResponse, err error) {
		return &slack.OAuthAccessResponse{TeamID: "TEID", WebhookURL: "http://a/a", ChannelID: "CHID", RedirectURL: redirectServer.URL}, nil
	}

	a := &API{db: storage.New(fm), slackClient: sm}
	server := httptest.NewServer(a.bootRouter())
	defer server.Close()

	resp, err := http.Get(server.URL + "/init?code=abc")
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}

	if string(b) != "redirected" {
		t.Error("Redirect not working")
		return
	}
}

func TestNotFound(t *testing.T) {
	fm := &mockFirebase{}
	sm := &mockSlack{}

	a := &API{db: storage.New(fm), slackClient: sm}
	server := httptest.NewServer(a.bootRouter())
	defer server.Close()

	resp, err := http.Get(server.URL + "/notfound")
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}

	if string(b) != `{"message":"not found"}` {
		t.Error("NotFound not working")
		return
	}
}
