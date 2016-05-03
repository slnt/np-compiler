package convert

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/slantin/np-compiler/noonpacific"
)

func getArtwork(npp *noonpacific.Playlist) ([]byte, error) {
	req, err := http.NewRequest("GET", npp.CoverURL, nil)
	if err != nil {
		return nil, err
	}
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

	return bytes, nil
}

func saveArtwork(id int, artwork []byte) error {
	fname := fmt.Sprintf("noon%dcover.jpg", id)
	os.Remove(fname)
	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	_, err = file.Write(artwork)
	if err != nil {
		return err
	}
	return nil
}
