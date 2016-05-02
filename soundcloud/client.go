package soundcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"gopkg.in/validator.v2"

	log "github.com/sirupsen/logrus"
)

const baseURL = "https://api.soundcloud.com"

// Auth is authentication data retrieved from the Soundcloud API
type Auth struct {
	Token string `json:"access_token" validate:"nonzero"`
	Scope string `json:"scope"`
}

// A Client is used to make requests to the Soundcloud API
type Client struct {
	url      string
	username string
	password string
	id       string
	secret   string
	auth     Auth
}

// NewClient creates a new client for the Soundcloud api
func NewClient(cfg Config) *Client {
	c := &Client{
		url:      baseURL,
		username: cfg.Username,
		password: cfg.Password,
		id:       cfg.Client.ID,
		secret:   cfg.Client.Secret,
	}

	return c
}

// GetAuthToken requests an oauth2 token from the api's /oauth2/token endpoint
func (c *Client) GetAuthToken() error {
	form := make(url.Values)
	form.Set("grant_type", "password")
	form.Set("client_id", c.id)
	form.Set("client_secret", c.secret)
	form.Set("username", c.username)
	form.Set("password", c.password)
	form.Set("scope", "*")

	res, err := http.PostForm(fmt.Sprintf("%s/oauth2/token", c.url), form)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var auth Auth
	if err := json.Unmarshal(resBody, &auth); err != nil {
		return err
	}

	if err := validator.Validate(auth); err != nil {
		return err
	}

	c.auth = auth
	return nil
}

type wrappedPlaylist struct {
	Playlist *Playlist `json:"playlist"`
}

func wrapPlaylist(playlist *Playlist) *wrappedPlaylist {
	return &wrappedPlaylist{Playlist: playlist}
}

// UploadPlaylist takes the given playlist and tries to upload it using the client credentials
func (c *Client) UploadPlaylist(playlist *Playlist) error {
	body, err := json.Marshal(wrapPlaylist(playlist))
	if err != nil {
		return err
	}

	query := make(url.Values)
	query.Set("format", "json")
	query.Set("oauth_token", c.auth.Token)
	query.Set("client_id", c.id)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/playlists?%s", c.url, query.Encode()),
		bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Add("ACCEPT", "application/json")

	// TODO: fix this shit v2, keep getting 422 unprocessable entity
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"status": res.Status,
			"body":   string(resBody),
		}).Error("Request failed")
		return fmt.Errorf("Failed to do request, got: %s", res.Status)
	}

	return nil
}
