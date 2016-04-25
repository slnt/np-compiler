package main

import (
	"flag"
	"log"

	"github.com/slantin/np-compiler/config"
	"github.com/slantin/np-compiler/convert"
	"github.com/slantin/np-compiler/noonpacific"
	"github.com/slantin/np-compiler/soundcloud"
)

var id = flag.Int("id", 1, "ID of playlist to upload")

func main() {
	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	npPlaylist, err := noonpacific.GetPlaylist(*id)
	if err != nil {
		log.Fatalf("Failed to get playlist data: %v", err)
	}

	client := soundcloud.NewClient(cfg.SoundCloud.ClientID, cfg.SoundCloud.ClientSecret)
	err = client.GetAuthToken()
	if err != nil {
		log.Fatalf("Failed to get SoundCloud auth token: %v", err)
	}

	scPlaylist, err := convert.NPtoSC(npPlaylist)
	if err != nil {
		log.Fatalf("Failed to convert playlist: %v", err)
	}

	err = client.UploadPlaylist(scPlaylist)
	if err != nil {
		log.Fatalf("Failed to upload playlist to SoundCloud: %v", err)
	}

}
