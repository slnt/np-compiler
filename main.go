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

	log.Info("Getting NoonPacific playlist data")
	npp, err := noonpacific.GetPlaylist(*id)
	if err != nil {
		log.Fatalf("Failed to get playlist data: %v", err)
	}

	log.Info("Getting playlist artwork")
	artwork, err := npp.GetArtwork(cfg.SaveArtwork)
	if err != nil {
		log.Fatalf("Failed to get playlist artwork: %v", err)
	}

	log.Info("Creating SoundCloud client and authorizing")
	client := soundcloud.NewClient(cfg.SoundCloud)
	err = client.GetAuthToken()
	if err != nil {
		log.Fatalf("Failed to get SoundCloud auth token: %v", err)
	}

	log.Info("Converting playlist")
	scp, err := convert.NPtoSC(npp)
	if err != nil {
		log.Fatalf("Failed to convert playlist: %v", err)
	}
	scp.ArtworkData = artwork

	log.Info("Uploading playlist to SoundCloud")
	res, err := client.UploadPlaylist(scp)
	if err != nil {
		log.Fatalf("Failed to upload playlist to SoundCloud: %v", err)
	}
	log.WithField("id", res.ID).Info("Uploaded playlist")

	// log.Info("Uploading playlist artwork to SoundCloud")
	// err = client.UploadArtwork(res.ID, artwork)
	// if err != nil {
	// 	log.Fatalf("Failed to upload playlist artwork: %v", err)
	// }

	log.Info("Successfully created playlist!")
}
