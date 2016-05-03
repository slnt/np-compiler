package noonpacific

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // TODO:fixit
	},
}

// NoonRegexp is the regular expressions for the title of Noon Pacific playlists
var NoonRegexp = regexp.MustCompile(`^NOON \/\/ \d+$`)

// Endpoint is the http API endpoint to get Noon Pacific track data
var Endpoint = "https://api.colormyx.com/v1/noon-pacific/playlists/%d/?detail=true"

// GetPlaylist hits Endpoint to get the playlist with the given ID. If no playlist
// exists, or if the playlist name does not match NoonRegexp, returns an error.
func GetPlaylist(id int) (*Playlist, error) {
	res, err := client.Get(fmt.Sprintf(Endpoint, id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var playlist Playlist
	err = json.Unmarshal(body, &playlist)
	if err != nil {
		return nil, err
	}

	if !NoonRegexp.MatchString(playlist.Name) {
		return nil, fmt.Errorf("Invalid playlist name: %v", playlist.Name)
	}

	return &playlist, nil
}

// GetArtwork tries to get the binary data of the playlists cover
func (npp *Playlist) GetArtwork(save bool) ([]byte, error) {
	req, err := http.NewRequest("GET", npp.CoverURL, nil)
	if err != nil {
		return nil, err
	}
	// TODO: :?
	// req.Header.Add("Accept", "image/jpeg")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if save {
		saveArtwork(npp.ID, bytes)
	}

	return bytes, nil
}

func saveArtwork(id int, artwork []byte) (*os.File, error) {
	fname := fmt.Sprintf("noon%dcover.jpg", id)
	os.Remove(fname)
	file, err := os.Create(fname)
	if err != nil {
		return nil, err
	}
	_, err = file.Write(artwork)
	if err != nil {
		return nil, err
	}
	return file, err
}
