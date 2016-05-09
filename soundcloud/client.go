package soundcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func (c *Client) newRequest(method, endpoint, format string, body io.Reader) (*http.Request, error) {
	query := make(url.Values)
	query.Set("format", format)
	query.Set("oauth_token", c.auth.Token)
	query.Set("client_id", c.id)

	return http.NewRequest(method,
		fmt.Sprintf("%s%s?%s", c.url, endpoint, query.Encode()),
		body)
}

func wrapPlaylist(playlist *Playlist) *wrappedPlaylist {
	return &wrappedPlaylist{Playlist: playlist}
}

// UploadPlaylist takes the given playlist and tries to upload it using the client credentials
func (c *Client) UploadPlaylist(playlist *Playlist) (*Playlist, error) {
	body, err := json.Marshal(wrapPlaylist(playlist))
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("POST", "/playlists", "json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Length", string(len(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		log.WithFields(log.Fields{
			"status": res.Status,
			"body":   string(resBody),
		}).Error("Request failed")
		return nil, fmt.Errorf("Failed to do request, got: %s", res.Status)
	}

	var rp Playlist
	err = json.Unmarshal(resBody, &rp)
	if err != nil {
		return nil, err
	}

	return &rp, nil
}

type artworkRequest struct {
	Playlist struct {
		ID          int    `json:"id"`
		ArtworkData []byte `json:"artwork_data"`
	} `json:"playlist"`
}

// UploadArtwork uploads the given artwork to the specified playlist
func (c *Client) UploadArtwork(playlistID int, artwork []byte) error {
	var areq artworkRequest
	areq.Playlist.ID = playlistID
	areq.Playlist.ArtworkData = artwork

	body, err := json.Marshal(&areq)
	if err != nil {
		return err
	}

	req, err := c.newRequest("PUT", "/playlists", "json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Length", string(len(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if !(res.StatusCode >= 200 && res.StatusCode < 300) {
		log.WithFields(log.Fields{
			"status": res.Status,
			// "body":   string(rgesBody),
		}).Error("Failed to upload artwork")
		return fmt.Errorf("Failed to do request, got: %s", res.Status)
	}

	// resBody, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return err
	// }
	//
	// var rp Playlist
	// err = json.Unmarshal(resBody, &rp)
	// if err != nil {
	// 	return err
	// }

	return nil
}
