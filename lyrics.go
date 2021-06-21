package swaglyrics_go

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/srevinsaju/swaglyrics-go/types"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var tag = regexp.MustCompile("(?s)<.*?>")

var httpClient = http.Client{
	Timeout: ApiTimeout * time.Second,
}

func crawlGeniusWebPage(urlData string) (*http.Response, error) {
	fetchUrl := fmt.Sprintf("https://genius.com/%s-lyrics", urlData)
	resp, err := httpClient.Get(fetchUrl)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getLyrics(song types.Song) (string, error) {
	urlData := Stripper(song)
	if strings.HasPrefix(urlData, "-") {
		return "", InvalidSongError
	}
	// format the url with the url path
	resp, err := crawlGeniusWebPage(urlData)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		form := url.Values{}
		form.Add("song", song.Track)
		form.Add("artist", song.Artist)

		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("%s/stripper", BackendUrl),
			strings.NewReader(form.Encode()),
		)
		if err != nil {
			return "", err
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, err = httpClient.Do(req)
		if err != nil {
			return "", err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		urlData = string(body[:])
		resp, err = crawlGeniusWebPage(urlData)
		if err != nil {
			return "", err
		}
	}

	geniusWebPageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	geniusWebPage := string(geniusWebPageBytes[:])
	geniusHtml := soup.HTMLParse(geniusWebPage)
	lyricsPath := geniusHtml.Find("div", "class", "lyrics")
	if lyricsPath.Error == nil {
		cleanedHtml := strings.Replace(lyricsPath.HTML(), "<br/>", "", -1)
		cleanedHtml = tag.ReplaceAllString(cleanedHtml, "")
		cleanedText := html.UnescapeString(cleanedHtml)
		return strings.Trim(cleanedText, "\n\r "), nil
	} else {
		var lyricsPaths []string
		lyricsPathDivs := geniusHtml.FindAll("div")

		for i := range lyricsPathDivs {

			if strings.HasPrefix(lyricsPathDivs[i].Attrs()["class"], "Lyrics__Container") {
				cleanedHtml := strings.Replace(lyricsPathDivs[i].HTML(), "<br/>", "\n", -1)
				cleanedHtml = tag.ReplaceAllString(cleanedHtml, "")
				cleanedText := html.UnescapeString(cleanedHtml)
				lyricsPaths = append(lyricsPaths, cleanedText)
			}
		}
		return strings.Join(lyricsPaths, "\n"), nil
	}
}

// GetLyrics fetches a song from the genius api
func GetLyrics(song types.Song) (string, error) {
	lyrics, err := getLyrics(song)
	if err != nil {
		return fmt.Sprintf("Couldn't get lyrics for %s by %s", song.Track, song.Artist), err
	}
	lyrics = strings.Trim(lyrics, "\n\r ")
	return lyrics, nil
}
