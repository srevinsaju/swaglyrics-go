package swaglyrics_go

import "errors"

var InvalidSongError = errors.New("invalid song")
var FetchLyricsError = errors.New("couldn't fetch lyrics from genius")
