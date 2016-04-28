package soundcloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const baseURL = "https://api.soundcloud.com"

type Auth struct {
	Token   string `json:"access_token"`
	Refresh string `json:"refresh_token"`
	Scope   string `json:"scope"`
	Expires int    `json:"expires_in"`
}

type Client struct {
	url      string
	username string
	password string
	id       string
	secret   string
	auth     Auth
}

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

func (c *Client) GetAuthToken() error {
	query := make(url.Values)
	query.Set("grant_type", "password")
	query.Set("client_id", c.id)
	query.Set("client_secret", c.secret)
	query.Set("username", c.username)
	query.Set("password", c.password)
	query.Set("scope", "")

	res, err := http.PostForm(fmt.Sprintf("%s/oauth2/token", c.url), query)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var auth Auth
	err = json.Unmarshal(resBody, &auth)
	if err != nil {
		return err
	}

	c.auth = auth
	return nil
}

func (c *Client) UploadPlaylist(playlist *Playlist) error {

	return nil
}
