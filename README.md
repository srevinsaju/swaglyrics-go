# swaglyrics-go

~~Look ma, I copied [that][swaglyrics] code!~~ Go lang rewrite of [swaglyrics][swaglyrics] ✨



From [Swaglyrics/Swaglyrics-for-Spotify][swaglyrics]'s README:

> Fetches the currently playing song from Spotify on Windows,
> Linux and macOS and displays the lyrics in the command-line,
> browser tab or in a desktop application. Refreshes automatically
> when song changes. The lyrics are fetched from Genius. Turns out
> Deezer already has this feature in-built but with swaglyrics, you
> can have it in Spotify as well.

This rewrite doesn't do all that, but, you can get the lyrics for a song
from your terminal: just do
```bash
sl --track 'bad guy' --artist 'Billie Eilish'
```
and you get the lyrics!

I'm mainly trying to build this project as far as I can,
~~for practice and to learn and work with more technologies and platforms~~ because
I had to use it in a Go project, and didn't want to add python dependencies
to it.

Initially ~~developed~~ copied this for personal use, and still is.
~~Pretty much~~ This is not functionality oriented, It Just Werkz™ --
I usually develop something that I can see helping me
and other users in the same situation. ~~Packaged so I can
first hand handle production-ready code to an extent and to make
distribution and usage easier.~~ Used Go lang, because it makes package distribution
very easy. Finally, I don't need to very about a million `pip` packages,
version conflicts, etc...


## Installation

```bash
sudo wget https://github.com/srevinsaju/swaglyrics-go/releases/download/continuous/sl -O /usr/local/bin/sl
sudo chmod +x /usr/local/bin/sl
sl --help
```


## Usage

`usage: sl --track 'Foo' --artist 'Bar'`

```
NAME:
   swaglyrics-go - Get the lyrics for a song

USAGE:
   sl [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --track value   Song title
   --artist value  Song artist
   --help, -h      show help (default: false)
```

Before using, you should check [USING.txt][using_txt] to comply with the Genius ToS.
There's a copy included inside the package as well.


## API

A simple broiler plate to use swaglyrics in your golang project:

```go
package main

import (
	"fmt"
	"log"

	sl "github.com/srevinsaju/swaglyrics-go"
	"github.com/srevinsaju/swaglyrics-go/types"
)

func main() {
	lyrics, err := sl.GetLyrics(types.Song{Track: "Butter", Artist: "BTS"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(lyrics)
}
```

## Building

```bash
git clone https://github.com/srevinsaju/swaglyrics-go
cd swaglyrics-go/cmd/sl
go build .
./sl --help
```


## References
* The original Python implementation: see [Swaglyrics-for-Spotify][swaglyrics].
* For its LICENSE, see [LICENSE.md][swaglyrics_license].
* [The Swaglyrics Project][swaglyrics_project]


## License
Since `swaglyrics-go` is a port of `swaglyrics` from python, respect its license.

See [LICENSE][license]


[using_txt]: https://github.com/srevinsaju/swaglyrics-go/blob/master/USING.txt
[swaglyrics]: https://github.com/Swaglyrics/Swaglyrics-for-Spotify
[swaglyrics_project]: https://github.com/Swaglyrics/
[swaglyrics_license]: https://github.com/SwagLyrics/SwagLyrics-For-Spotify/blob/master/LICENSE.md
[license]: https://github.com/srevinsaju/swaglyrics-go/blob/master/LICENSE


