package swaglyrics_go

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
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
var br = regexp.MustCompile("<br/>")

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
	if resp.StatusCode != 200 && strings.Contains(song.Artist, ",") {
		// has multiple supporting artist
		// let's try again with only the first artist
		return getLyrics(types.Song{
			Track:  song.Track,
			Artist: strings.Split(song.Artist, ",")[0],
		})
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	var lyricsContainerHtml string
	doc.Find(".lyrics").First().Each(func(i int, selection *goquery.Selection) {
		lyricsContainerHtml = selection.Text()
	})
	if lyricsContainerHtml == "" {

		selector := doc.Find("div[class^='Lyrics__Container'], div[class*=' Lyrics__Container']")
		lyricsBuilder := strings.Builder{}
		selector.Each(func(i int, selection *goquery.Selection) {
			data, err := selection.Html()
			if err != nil {
				return
			}
			cleanedHtml := br.ReplaceAllString(data, "\n")
			cleanedHtml = tag.ReplaceAllString(cleanedHtml, "")

			cleanedHtml = html.UnescapeString(cleanedHtml)
			lyricsBuilder.WriteString(cleanedHtml)
			lyricsBuilder.WriteString("\n")
		})
		lyricsContainerHtml = lyricsBuilder.String()
	}
	return lyricsContainerHtml, nil
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
