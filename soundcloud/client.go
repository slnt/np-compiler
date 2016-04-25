package soundcloud

type Client struct {
	url         string
	id          string
	secret      string
	oauth2token string
}

func NewClient(id, secret string) *Client {
	c := &Client{
		url:    "https://api.soundcloud.com",
		id:     id,
		secret: secret,
	}

	return c
}

func (c *Client) GetAuthToken() error {

	return nil
}

func (c *Client) UploadPlaylist(playlist *Playlist) error {

	return nil
}
