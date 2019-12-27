package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-soundcloud"
)

func main() {
	client := soundcloud.NewClient()
	res, _ := client.Track.List()
	tracks, _ := json.Marshal(res)
	fmt.Println(string(tracks))
}
