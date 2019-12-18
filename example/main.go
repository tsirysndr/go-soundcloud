package main

import (
	"fmt"

	"github.com/tsirysndr/go-soundcloud"
)

func main() {
	client := soundcloud.NewClient()
	track, _ := client.Track.Get("13158665")
	tracks, err := client.Track.List()
	fmt.Println(track)
	fmt.Println(tracks, err)
}
