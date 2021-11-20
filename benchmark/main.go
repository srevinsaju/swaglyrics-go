package main

import (
	"encoding/json"
	"fmt"
	sl "github.com/srevinsaju/swaglyrics-go"
	"github.com/srevinsaju/swaglyrics-go/types"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var errored = 0
var totalTime time.Duration = 0

type ArtistMeta struct {
	Name string `json:"name"`
}

type TrackMeta struct {
	Name    string       `json:"name"`
	Artists []ArtistMeta `json:"artists"`
}

type Item struct {
	Track TrackMeta `json:"track"`
}

type TopUS50 struct {
	Items []Item `json:"items"`
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
	totalTime += elapsed
}

func GetLyrics(song types.Song, meta string) (string, error) {
	defer timeTrack(time.Now(), meta)
	lyrics, err := sl.GetLyrics(song)
	return lyrics, err
}

func main() {

	jsonFile := os.Args[len(os.Args)-1]
	// read our opened jsonFile as a byte array.
	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	topus50 := &TopUS50{}

	err = json.Unmarshal(data, topus50)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range topus50.Items {
		track := topus50.Items[i].Track
		name := track.Name
		artist := track.Artists[0].Name
		lyrics, err := GetLyrics(types.Song{Track: name, Artist: artist}, fmt.Sprintf("%s by %s", name, artist))
		if err != nil || strings.HasPrefix(lyrics, "Couldn't get lyrics for") {
			fmt.Printf("Failed to get lyrics for %s by %s\n", name, artist)
			errored += 1
		}
	}

	fmt.Printf(
		"Successful for {%d}/50 cases.\n"+
			"Total time {%s}.\n"+
			"Avg. time {%s/50}s.\n",
		50-errored,
		totalTime,
		totalTime/50,
	)
}
