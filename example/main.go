package main

import (
	"fmt"

	"github.com/tsirysndr/go-soundcloud"
)

func main() {
	client := soundcloud.NewClient()
	track, _ := client.Track.Get("13158665")
	tracks, _ := client.Track.List()
	playlist, _ := client.Playlist.Get("405726")
	playlists, _ := client.Playlist.List()
	fmt.Println(track)
	fmt.Println(tracks)
	fmt.Println(playlist)
	fmt.Println(playlists)
}
