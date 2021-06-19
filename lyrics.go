package swaglyrics_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/srevinsaju/swaglyrics-go/types"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var tag = regexp.MustCompile("(?s)<.*?>")

var httpClient = http.Client{
	Timeout: API_TIMEOUT * time.Second,
}

func crawlGeniusWebPage(urlData string) (*http.Response, error) {
	url := fmt.Sprintf("https://genius.com/%s-lyrics", urlData)
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetLyrics fetches a song from the genius api
func GetLyrics(song types.Song) (string, error) {
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
		postBody, _ := json.Marshal(map[string]string{
			"song":   song.Track,
			"artist": song.Artist,
		})
		responseBody := bytes.NewBuffer(postBody)

		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("%s/stripper", BACKEND_URL),
			responseBody,
		)
		if err != nil {
			return "", err
		}
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
