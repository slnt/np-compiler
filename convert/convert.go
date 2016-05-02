package convert

import (
	"fmt"
	"time"

	"github.com/slantin/np-compiler/noonpacific"
	"github.com/slantin/np-compiler/soundcloud"
)

var descriptionFmt = "http://noonpacific.com/#/mix/%d\n"

// NPtoSC takes a Noon Pacific playlist and convets it into a SoundCloud playlist
func NPtoSC(npp *noonpacific.Playlist) (*soundcloud.Playlist, error) {
	var scp soundcloud.Playlist

	release, err := time.ParseInLocation(time.RFC3339, npp.ReleaseDate, time.FixedZone("PDT", 0))
	if err != nil {
		return nil, err
	}

	scp.Title = npp.Name
	scp.Sharing = "public"
	scp.ReleasyYear = release.Year()
	scp.ReleaseMonth = int(release.Month())
	scp.ReleaseDay = release.Day()
	scp.ArtworkURL = npp.CoverURL
	scp.Type = "compilation"
	scp.Tags = "noonpacific"
	scp.Genre = "noonpacific"
	scp.Description = fmt.Sprintf(descriptionFmt, npp.ID)

	for _, track := range npp.Tracks {
		scp.Tracks = append(scp.Tracks, soundcloud.Track{
			ID: track.SoundCloudID,
		})
	}

	return &scp, nil
}
