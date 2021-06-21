package main

import (
	"bufio"
	"fmt"
	sl "github.com/srevinsaju/swaglyrics-go"
	"github.com/srevinsaju/swaglyrics-go/types"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

func input(r *bufio.Reader, msg string) (res string) {
	fmt.Print(msg)
	res, _ = r.ReadString('\n')
	res = strings.Trim(res, "\r\n")
	return
}

func main() {
	app := &cli.App{
		Name: "swaglyrics-go",
		Usage: "Get the lyrics for a song",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "track", Usage: "Song title"},
			&cli.StringFlag{Name: "artist", Usage: "Song artist"},
		},
		Action: func(c *cli.Context) error {

			track := c.String("track")
			artist := c.String("artist")
			if track == "" || artist == "" {
				log.Fatalln("Please provide the track and artist names.")
				return nil
			}
			fmt.Printf("Getting lyrics for %s by %s\n", track, artist)
			lyrics, err := sl.GetLyrics(types.Song{
				Track:  track,
				Artist: artist,
			})
			if err != nil {
				return err
			}
			lyrics = strings.Trim(lyrics, "\n\r ")
			if lyrics == "" {
				fmt.Printf("Couldn't get lyrics for %s by %s\n", track, artist)
			} else {
				fmt.Println(lyrics)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
