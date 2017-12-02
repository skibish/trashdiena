package slack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOAuthAccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"ok":true,"team_id":"TM12","incoming_webhook":{"channel_id":"CH12","configuration_url":"http://confurl","url":"http://redirecturl"}}`)
	}))
	defer ts.Close()

	sl := &Slack{url: ts.URL + "/", hc: &http.Client{}}

	resp, err := sl.OAuthAccess("hgf")
	if err != nil {
		t.Error(err)
		return
	}

	if resp.TeamID != "TM12" {
		t.Errorf("Expected teamID=%s, got %s", resp.TeamID, "TM12")
	}

	if resp.ChannelID != "CH12" {
		t.Errorf("Expected channelID=%s, got %s", resp.ChannelID, "CH12")
	}

	if resp.WebhookURL != "http://redirecturl" {
		t.Errorf("Expected webhookURL=%s, got %s", resp.WebhookURL, "http://redirecturl")
	}

	if resp.RedirectURL != "http://confurl" {
		t.Errorf("Expected redirectURL=%s, got %s", resp.RedirectURL, "http://confurl")
	}
}

func TestOAuthAccessError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"ok":false,"error":"some error occured"}`)
	}))
	defer ts.Close()

	sl := &Slack{url: ts.URL + "/", hc: &http.Client{}}

	_, err := sl.OAuthAccess("hgf")
	if err.Error() != "some error occured" {
		t.Error("Should be error. but everything is fine")
		return
	}
}

func TestOAuthAccessParseError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{aaa`)
	}))
	defer ts.Close()

	sl := &Slack{url: ts.URL + "/", hc: &http.Client{}}

	_, err := sl.OAuthAccess("hgf")
	if err.Error() != "invalid character 'a' looking for beginning of object key string" {
		t.Error("Should be error. but everything is fine")
		return
	}
}
