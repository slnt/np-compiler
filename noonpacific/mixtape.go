package noonpacific

// Example JSON
// {
//     "id":4722,
//     "mixtape":565,
//     "title":"True Colours (Hannes Fischer Vocal Remix)",
//     "artist":"Radio Atlantis",
//     "slug":"radio-atlantis-true-colours-hannes-fischer-vocal-r",
//     "streamable":true,
//     "duration":378157,
//     "external_id":285412939,
//     "stream_url":"https://api.soundcloud.com/tracks/285412939/stream",
//     "permalink_url":"http://soundcloud.com/hannes-fischer/radio-atlantis-true-colours-hannes-fischer-vocal-remix",
//     "artwork_url":"https://i1.sndcdn.com/artworks-000186007735-qa5l7i-large.jpg",
//     "purchase_url":"https://itunes.apple.com/de/album/true-colours-hannes-fischer/id1152996015",
//     "download_url":"",
//     "ticket_url":"",
//     "play_count":0,
//     "order":1
// }

// Track is an element of a Noon Pacific Playlist
type Track struct {
	ID         int    `json:"external_id"`
	Mixtape    int    `json:"mixtape"`
	Title      string `json:"title"`
	Artist     string `json:"artist"`
	ArtworkURL string `json:"artwork_url"`
}

// A Tracklist is the tracklist for a mixtape
type Tracklist struct {
	Tracks []Track `json:"results"`
}

// Mixtape is a mixtape from the noonpacific api
type Mixtape struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Slug        string  `json:"slug"`
	Description string  `json:"description"`
	ArtworkURL  string  `json:"artwork_url"`
	Artist      string  `json:"artwork_credit"`
	ArtistURL   string  `json:"artwork_credit_url"`
	ReleaseTime string  `json:"release"`
	Tracks      []Track // populate later
}

// Mixtapes contains a list of mixtapes, the latest being the first
type Mixtapes struct {
	List []Mixtape `json:"results"`
}
