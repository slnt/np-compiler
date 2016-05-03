package soundcloud

// A Track on Soundcloud
type Track struct {
	ID int `json:"id"`
}

// A Playlist on Soundcloud
type Playlist struct {
	Title        string  `json:"title"`
	Sharing      string  `json:"sharing"`
	Tracks       []Track `json:"tracks"`
	Created      string  `json:"created_at"`
	ReleaseDay   int     `json:"release_day"`
	ReleaseMonth int     `json:"release_month"`
	ReleasyYear  int     `json:"release_year"`
	ArtworkURL   string  `json:"artwork_url"`
	Type         string  `json:"playlist_type"`
	Genre        string  `json:"genre"`
	Description  string  `json:"description"`
	Tags         string  `json:"tag_list"`
}
