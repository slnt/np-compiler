package main

import (
	"flag"

	log "github.com/sirupsen/logrus"

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

	npp, err := noonpacific.GetPlaylist(*id)
	if err != nil {
		log.Fatalf("Failed to get playlist data: %v", err)
	}

	client := soundcloud.NewClient(cfg.SoundCloud)
	err = client.GetAuthToken()
	if err != nil {
		log.Fatalf("Failed to get SoundCloud auth token: %v", err)
	}

	scp, err := convert.NPtoSC(npp)
	if err != nil {
		log.Fatalf("Failed to convert playlist: %v", err)
	}

	err = client.UploadPlaylist(scp)
	if err != nil {
		log.Fatalf("Failed to upload playlist to SoundCloud: %v", err)
	}
}
