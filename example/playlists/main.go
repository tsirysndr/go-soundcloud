package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-soundcloud"
)

func main() {
	client := soundcloud.NewClient()
	res, _ := client.Playlist.List()
	playlists, _ := json.Marshal(res)
	fmt.Println(string(playlists))
}
