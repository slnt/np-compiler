package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/slantin/np-compiler/config"
	"github.com/slantin/np-compiler/convert"
	"github.com/slantin/np-compiler/noonpacific"
	"github.com/slantin/np-compiler/print"
	"github.com/slantin/np-compiler/soundcloud"
)

func main() {
	cfg := config.Get()
	noonpacific.Init(cfg.NoonPacific.ClientID)

	log.Info("Getting NoonPacific playlist data")
	mixtape, err := noonpacific.LatestMixtape()
	if err != nil {
		log.Fatalf("Failed to get mixtape data: %v", err)
	}
	print.JSON(mixtape)

	log.Info("Getting playlist artwork")
	// artwork, err := mixtape.GetArtwork(cfg.SaveArtwork)
	_, err = mixtape.GetArtwork(cfg.SaveArtwork)
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
	playlist, err := convert.NPtoSC(mixtape)
	if err != nil {
		log.Fatalf("Failed to convert playlist: %v", err)
	}
	// playlist.ArtworkData = artwork

	log.Info("Uploading playlist to SoundCloud")
	res, err := client.UploadPlaylist(playlist)
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
