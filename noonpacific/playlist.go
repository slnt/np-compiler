package noonpacific

// Track is an element of a Noon Pacific Playlist
type Track struct {
	ID                     int    `json:"id"`
	Title                  string `json:"title"`
	ArtistDescription      string `json:"artist_description"`
	TrackID                int    `json:"track_id"`
	PlaylistID             int    `json:"playlist_id"`
	TrackNumber            int    `json:"track_number"`
	SoundCloudID           int    `json:"soundcloud_id"`
	SoundCloudPermalinkURL string `json:"soundcloud_permalink_url"`
	SoundCloudStreamURL    string `json:"soundcloud_stream_url"`
	Duration               int    `json:"duration"`
	SourceNotFound         bool   `json:"source_not_found"`
}

// Playlist is a Noon Pacific Playlist
type Playlist struct {
	Name             string  `json:"name"`
	ID               int     `json:"id"`
	Description      string  `json:"description"`
	PlaylistNumber   int     `json:"playlist_number"`
	TrackCount       int     `json:"track_count"`
	Tracks           []Track `json:"tracks"`
	ReleaseDate      string  `json:"release_date"`
	DateCreated      string  `json:"date_created"`
	CoverURL         string  `json:"cover_large"`
	CoverDescription string  `json:"cover_description"`
}
