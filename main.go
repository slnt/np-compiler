package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/slantin/np-compiler/noonpacific"
)

var id = flag.Int("id", 1, "ID of playlist to upload")

func main() {
	flag.Parse()

	playlist, err := noonpacific.GetPlaylist(*id)
	if err != nil {
		log.Fatalf("Failed to get playlist data: %v", err)
	}

	fmt.Println(*playlist)
}
