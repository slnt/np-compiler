package soundcloud

type Track struct {
	ID int `json:"id"`
}

type Playlist struct {
	Title        string  `json:"title"`
	Sharing      string  `json:"sharing"`
	Tracks       []Track `json:"tracks"`
	ReleaseDay   int     `json:"release_day"`
	ReleaseMonth int     `json:"release_month"`
	ReleasyYear  int     `json:"release_year"`
	ArtworkURL   string  `json:"artwork_url"`
	Type         string  `json:"playlist_type"`
}
