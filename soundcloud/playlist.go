package soundcloud

// A Track on Soundcloud
type Track struct {
	ID int `json:"id"`
}

// A Playlist on Soundcloud
type Playlist struct {
	ID           int     `json:"id,omitempty"`
	Title        string  `json:"title,omitempty"`
	Sharing      string  `json:"sharing,omitempty"`
	Tracks       []Track `json:"tracks,omitempty"`
	Created      string  `json:"created_at,omitempty"`
	ReleaseDay   int     `json:"release_day,omitempty"`
	ReleaseMonth int     `json:"release_month,omitempty"`
	ReleaseYear  int     `json:"release_year,omitempty"`
	ArtworkData  []byte  `json:"artwork_data,omitempty"`
	Type         string  `json:"playlist_type,omitempty"`
	Genre        string  `json:"genre,omitempty"`
	Description  string  `json:"description,omitempty"`
	Tags         string  `json:"tag_list,omitempty"`
}
