package noonpacific

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

// NoonRegexp is the regular expressions for the title of Noon Pacific playlists
var NoonRegexp = regexp.MustCompile(`^NOON \/\/ \d+$`)

// API is the api URL for the noonpacific api
var API = "https://beta.whitelabel.cool/api"

// ClientID is the magical client id gleaned from the api calls made by noonpacific.com
var ClientID string

var client = http.DefaultClient

// Init sets the clientID for requests
func Init(clientID string) {
	ClientID = clientID
}

// NewRequest creates a new http request to the given endpoint for the noonpacific api
func NewRequest(endpoint string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", API, endpoint), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Client", ClientID)

	return req, err
}

// LatestMixtape gets the latest noonpacific mixtape and returns it
func LatestMixtape() (*Mixtape, error) {
	req, err := NewRequest("/mixtapes/")
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var mixtapes Mixtapes
	err = json.Unmarshal(body, &mixtapes)
	if err != nil {
		return nil, err
	}
	if len(mixtapes.List) < 0 {
		return nil, fmt.Errorf("")
	}

	mixtape := &mixtapes.List[0]
	if !NoonRegexp.MatchString(mixtape.Title) {
		return nil, fmt.Errorf("Invalid mixtape title: %s", mixtape.Title)
	}

	if err = mixtape.getTracklist(); err != nil {
		return nil, err
	}

	return mixtape, nil
}

func (m *Mixtape) getTracklist() error {
	req, err := NewRequest(fmt.Sprintf("/tracks/?mixtape=%s", m.Slug))
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var tracklist Tracklist
	if err = json.Unmarshal(body, &tracklist); err != nil {
		return err
	}

	m.Tracks = tracklist.Tracks

	return nil
}

// GetArtwork tries to get the binary data of the playlists cover
func (m *Mixtape) GetArtwork(save bool) ([]byte, error) {
	req, err := http.NewRequest("GET", m.ArtworkURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "image/jpeg")

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
		saveArtwork(m.ID, bytes)
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
