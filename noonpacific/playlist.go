package noonpacific

var endpoint = "https://api.colormyx.com/v1/noon-pacific/playlists/%d/?detail=true"

type Track struct {
	ID                     int    `json:"id"`
	Title                  string `json:"title"`
	ArtistDescription      string `json:"artist_description"`
	SoundCloudID           int    `json:"soundcloud_id"`
	SoundCloudPermalinkURL string `json:"soundcloud_permalink_url"`
	SoundCloudStreamURL    string `json:"soundcloud_stream_url"`
	Duration               int    `json:"duration"`
	SourceNotFound         bool   `json:"source_not_found"`
	TrackID                int    `json:"track_id"`
	PlaylistID             int    `json:"playlist_id"`
	TrackNumber            int    `json:"track_number"`
}

type Playlist struct {
	Name              string  `json:"name"`
	ID                int     `json:"id"`
	Description       string  `json:"description"`
	TrackCount        int     `json:"track_count"`
	Tracks            []Track `json:"tracks"`
	ReleaseDate       string  `json:"release_date"`
	DateCreated       string  `json:"date_created"`
	Released          bool    `json:"released"`
	PlaylistNumber    int     `json:"playlist_number"`
	CoverURL          string  `json:"cover_large"`
	ReleaseWeekNumber int     `json:"release_week_number"`
	ReleaseYearNumber int     `json:"release_year_number"`
	CoverDescription  string  `json:"cover_description"`
	ReleaseWeek       string  `json:"release_week"`
}
