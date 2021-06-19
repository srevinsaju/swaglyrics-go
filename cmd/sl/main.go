package main

import (
	"fmt"
	sl "github.com/srevinsaju/swaglyrics-go"
	"github.com/srevinsaju/swaglyrics-go/types"
	"log"
)

func main() {
	lyrics, err := sl.GetLyrics(types.Song{
		Track:  "Bad Guy",
		Artist: "Billie Eilish",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(lyrics)
}
