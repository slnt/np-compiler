package soundcloud

type Client struct {
	url         string
	id          string
	secret      string
	password    string
	oauth2token string
}

func NewClient(cfg Config) *Client {
	c := &Client{
		url:      "https://api.soundcloud.com",
		id:       cfg.ClientID,
		secret:   cfg.ClientSecret,
		password: cfg.Password,
	}

	return c
}

func (c *Client) GetAuthToken() error {

	return nil
}

func (c *Client) UploadPlaylist(playlist *Playlist) error {

	return nil
}
