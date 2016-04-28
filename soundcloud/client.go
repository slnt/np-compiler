package soundcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"gopkg.in/validator.v2"
)

const baseURL = "https://api.soundcloud.com"

type Auth struct {
	Token string `json:"access_token" validate:"nonzero"`
	Scope string `json:"scope"`
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

func (c *Client) UploadPlaylist(playlist *Playlist) error {
	body, err := json.Marshal(*playlist)
	if err != nil {
		return err
	}

	query := make(url.Values)
	query.Set("access_token", c.auth.Token)
	query.Set("client_id", c.id)
	query.Set("client_secret", c.secret)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/playlists?%s", c.url, query.Encode()),
		bytes.NewReader(body))
	if err != nil {
		return err
	}

	// TODO: fix this shit, keep getting 401 unauthorized...
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	fmt.Println(res.Status)

	return nil
}
