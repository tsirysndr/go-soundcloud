package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-soundcloud"
)

func main() {
	client := soundcloud.NewClient()
	res, _ := client.Track.Get("13158665")
	track, _ := json.Marshal(res)
	fmt.Println(string(track))
}
