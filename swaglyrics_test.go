package swaglyrics_go

import (
	"github.com/srevinsaju/swaglyrics-go/types"
	"testing"
)

func stripper(track string, artist string) string {
	return Stripper(types.Song{Track: track, Artist: artist})
}

func assertEqual(fxOutput string, realOutput string) bool {
	return fxOutput == realOutput
}

func TestStripper(t *testing.T) {
	if matches := assertEqual(stripper("River (feat. Ed Sheeran)", "Eminem"), "Eminem-River"); !matches { t.Error() }
	if matches := assertEqual(stripper("Ain't My Fault - R3hab Remix", "Zara Larsson"), "Zara-Larsson-Aint-My-Fault"); !matches { t.Error() }
	if matches := assertEqual(stripper("1800-273-8255", "Logic"), "Logic-1800-273-8255"); !matches { t.Error() }
	if matches := assertEqual(stripper("Garota", "Erlend Øye"), "Erlend-ye-Garota"); !matches { t.Error() }
	if matches := assertEqual(stripper("Scream & Shout", "will.i.am"), "william-Scream-and-Shout"); !matches { t.Error() }
	if matches := assertEqual(stripper("Heebiejeebies - Bonus", "Aminé"), "Amine-Heebiejeebies"); !matches { t.Error() }
	if matches := assertEqual(stripper("FRÜHLING IN PARIS", "Rammstein"), "Rammstein-FRUHLING-IN-PARIS"); !matches { t.Error() }
	if matches := assertEqual(stripper("Chanel (Go Get It) [feat. Gunna & Lil Baby]", "Young Thug"), "Young-Thug-Chanel-Go-Get-It"); !matches { t.Error() }
	if matches := assertEqual(stripper("MONOPOLY (with Victoria Monét)", "Ariana Grande"), "Ariana-Grande-and-Victoria-Monet-MONOPOLY"); !matches { t.Error() }
	if matches := assertEqual(stripper("Seasons (with Sjava & Reason)", "Mozzy"), "Mozzy-Sjava-and-Reason-Seasons"); !matches { t.Error() }
	if matches := assertEqual(stripper("거품 안 넘치게 따라줘 [Life Is Good] (feat. Crush, Dj Friz)", "Dynamic Duo"), "Dynamic-Duo-Life-Is-Good"); !matches { t.Error() }
	if matches := assertEqual(stripper("Ice Hotel (ft. SZA)", "XXXTENTACION"), "XXXTENTACION-Ice-Hotel"); !matches { t.Error() }
	if matches := assertEqual(stripper("Zikr (From 'Amavas')", "Armaan Malik"), "Armaan-Malik-Zikr"); !matches { t.Error() }
}