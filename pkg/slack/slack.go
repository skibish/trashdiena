package slack

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// Slack describes slack struct
type Slack struct {
	url          string
	hc           *http.Client
	clientID     string
	clientSecret string
	redirectURL  string
}

// OAuthAccessAPIResponse is a respone for OAuthAccess request
type OAuthAccessAPIResponse struct {
	Ok              bool   `json:"ok"`
	AccessToken     string `json:"access_token"`
	Scope           string `json:"scope"`
	UserID          string `json:"user_id"`
	TeamName        string `json:"team_name"`
	TeamID          string `json:"team_id"`
	IncomingWebhook struct {
		Channel          string `json:"channel"`
		ChannelID        string `json:"channel_id"`
		ConfigurationURL string `json:"configuration_url"`
		URL              string `json:"url"`
	} `json:"incoming_webhook"`
}

// OAuthAccessResponse is an response of OAuthAccess method
type OAuthAccessResponse struct {
	WebhookURL  string
	ChannelID   string
	TeamID      string
	RedirectURL string
}

// New return new Slack instance
func New(clientID, clientSecret, redirectURL string) *Slack {
	return &Slack{
		url:          "https://slack.com/api/",
		hc:           &http.Client{},
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURL:  redirectURL,
	}
}

func (s *Slack) setupForm() *url.Values {
	return &url.Values{
		"url":           {s.redirectURL},
		"client_id":     {s.clientID},
		"client_secret": {s.clientSecret},
		"redirect_url":  {s.redirectURL},
	}
}

// OAuthAccess return OAuth data
func (s *Slack) OAuthAccess(code string) (r *OAuthAccessResponse, err error) {
	u := s.url + "oauth.access"

	form := s.setupForm()
	form.Add("code", code)

	req, err := http.NewRequest(http.MethodPost, u, strings.NewReader(form.Encode()))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := s.hc.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	var jsonResponse OAuthAccessAPIResponse
	err = json.NewDecoder(res.Body).Decode(&jsonResponse)
	if err != nil {
		return
	}

	r = &OAuthAccessResponse{
		WebhookURL:  jsonResponse.IncomingWebhook.URL,
		ChannelID:   jsonResponse.IncomingWebhook.ChannelID,
		TeamID:      jsonResponse.TeamID,
		RedirectURL: jsonResponse.IncomingWebhook.ConfigurationURL,
	}

	return
}
