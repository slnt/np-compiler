package noonpacific

import (
	"fmt"
	"ioutil"
	"net/http"
)

func Get(id int) (*Playlist, error) {

	var playlist Playlist
	res, err := http.DefaultClient.Get(fmt.Sprintf(endpoint, id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &playlist, nil
}
