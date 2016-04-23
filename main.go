package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/slantin/np-compiler/noonpacific"
)

var id = flag.Int("id", 1, "")

func main() {

	playlist, err := noonpacific.Get(*id)
	// var playlist noonpacific.Playlist
	res, err := http.DefaultClient.Get(fmt.Sprintf(noonpacific.Endpoint, id))
	if err != nil {
		log.Fatalf("Could not get data for Noon Pacific Playlist: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.Readall(res.Body)
	if err != nil {
		return
	}
}
