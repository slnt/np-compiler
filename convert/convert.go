package convert

import (
	"fmt"
	"time"

	"github.com/slantin/np-compiler/noonpacific"
	"github.com/slantin/np-compiler/soundcloud"
)

var (
	descriptionFmt = "http://noonpacific.com/#/mix/%s\n"
	scTimeFmt      = "2006/01/02 15:04:05 -0700"
)

// NPtoSC takes a Noon Pacific playlist and convets it into a SoundCloud playlist
func NPtoSC(m *noonpacific.Mixtape) (*soundcloud.Playlist, error) {
	release, err := time.ParseInLocation(time.RFC3339, m.ReleaseTime, time.FixedZone("PDT", 0))
	if err != nil {
		return nil, err
	}

	scp := &soundcloud.Playlist{
		Title:        m.Title,
		Sharing:      "public",
		Created:      release.Format(scTimeFmt),
		ReleaseYear:  release.Year(),
		ReleaseMonth: int(release.Month()),
		ReleaseDay:   release.Day(),
		Type:         "compilation",
		Tags:         "noonpacific",
		Genre:        "noonpacific",
		Description:  fmt.Sprintf(descriptionFmt, m.Slug),
	}

	for _, track := range m.Tracks {
		scp.Tracks = append(scp.Tracks, soundcloud.Track{ID: track.ID})
	}

	return scp, nil
}
