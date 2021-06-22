package swaglyrics_go

import (
	"fmt"
	"github.com/mozillazg/go-unidecode"
	"github.com/srevinsaju/swaglyrics-go/types"
	"regexp"
	"strings"
)

// https://github.com/SwagLyrics/SwagLyrics-For-Spotify/blob/99fe764a9e45cac6cb9fcdf724c7d2f8cb4524fb/swaglyrics/cli.py#L18-L24
// matches braces with feat included or text after -, also adds support for Bollywood songs by matching (From "<words>")

var Brc = regexp.MustCompile(`([(\[](feat|ft|From|Feat|from "[^"]*")[^)\]]*[)\]]|- .*)`)
var Aln = regexp.MustCompile(`[^ \-a-zA-Z0-9]+`)               // matches non space or - or alphanumeric characters
var Spc = regexp.MustCompile(` *- *| +`)                       // matches one or more spaces
var Wth = regexp.MustCompile(` *\(with ([^)]+)\)`)             // capture text after with
var Nlt = regexp.MustCompile(`[^\x00-\x7F\x80-\xFF\p{Latin}]`) // match only latin characters,
//var Nlt = regexp.MustCompile(`[]]`)
var unsafeChar = []string{"/", "_", "!"}
var someChar = []string{"Ø", "ø"}

// built using latin character tables (basic, supplement, extended a,b and extended additional)

// Stripper Generate the url path given the song and artist to format the Genius URL with.
// Strips the song and artist of special characters and unresolved text such as 'feat.' or text within braces.
// Then concatenates both with hyphens replacing the blank spaces.
func Stripper(song types.Song) string {
	artist := song.Artist
	track := song.Track
	track = Brc.ReplaceAllString(song.Track, "")
	track = strings.Trim(track, "\n\r ")
	ft := Wth.FindStringSubmatch(track)
	if len(ft) != 0 {
		track = strings.Replace(track, ft[0], "", -1)
		supportingArtists := ft[1]
		if strings.Contains(supportingArtists, "&") {
			artist = fmt.Sprintf("%s-%s", artist, supportingArtists)
		} else {
			artist = fmt.Sprintf("%s-and-%s", artist, supportingArtists)
		}
	}

	songData := fmt.Sprintf("%s-%s", artist, track)
	urlData := strings.Replace(songData, "&", "and", -1)

	for i := range unsafeChar {
		urlData = strings.Replace(urlData, unsafeChar[i], " ", -1)
	}
	for i := range someChar {
		urlData = strings.Replace(urlData, someChar[i], "", -1)
	}

	urlData = Nlt.ReplaceAllString(urlData, "")
	urlData = unidecode.Unidecode(urlData)
	urlData = Aln.ReplaceAllString(urlData, "")
	urlData = Spc.ReplaceAllString(urlData, "-")
	return urlData

}


func NormalizeArtist(artist string) string {

	urlData := strings.Replace(artist, "&", ", ", -1)


	for i := range unsafeChar {
		urlData = strings.Replace(urlData, unsafeChar[i], " ", -1)
	}
	for i := range someChar {
		urlData = strings.Replace(urlData, someChar[i], "", -1)
	}

	urlData = Nlt.ReplaceAllString(urlData, "")
	urlData = unidecode.Unidecode(urlData)
	urlData = Aln.ReplaceAllString(urlData, "")
	urlData = Spc.ReplaceAllString(urlData, " ")
	urlData = strings.Trim(urlData, " ")
	if urlData == "" {
		return artist
	}
	return urlData

}
