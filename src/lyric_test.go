package src

import (
	"fmt"
	"testing"
)

var (
	songTest = Song{Lyric: []string{"the horse and the hound and the horn that belonged to",
		"the farmer sowing his corn that kept",
		"the rooster that crowed in the morn that woke",
		"the priest all shaven and shorn that married",
		"the man all tattered and torn that kissed",
		"the maiden all forlorn that milked",
		"the cow with the crumpled horn that tossed",
		"the dog that worried",
		"the cat that killed",
		"the rat that ate",
		"the malt that lay in",
		"the house that Jack built",
	}}
	getOneLyricActual  string = "the house that Jack built"
	getOneObjectActual string = "This is the house"
	getMoreThanSongLine string = fmt.Sprintf("Input more than %d song line", len(songTest.Lyric))
	getZeroThanSongLine string = "Input cannot zero"
)

func TestGetOneLyric(t *testing.T) {
	t.Logf("Get one lyric : %s", getOneLyricActual)
	if songTest.Recite(1) != getOneLyricActual {
		t.Errorf("WRONG! it must %s", getOneLyricActual)
	}
}

func TestGetOneObject(t *testing.T) {
	t.Logf("Get one object : %s", getOneObjectActual)
	if songTest.ReciteObject(1) != getOneObjectActual {
		t.Errorf("WRONG! it must %s", getOneObjectActual)
	}
}

func TestGetLyricMoreThanSongLine(t *testing.T) {
	t.Logf("Get lyric more than song line : %s", getMoreThanSongLine)
	if songTest.Recite(len(songTest.Lyric)+1) != getMoreThanSongLine {
		t.Errorf("WRONG! it must %s", getMoreThanSongLine)
	}
}

func TestGetObjectMoreThanSongLine(t *testing.T) {
	t.Logf("Get object more than song line : %s", getMoreThanSongLine)
	if songTest.ReciteObject(len(songTest.Lyric)+1) != getMoreThanSongLine {
		t.Errorf("WRONG! it must %s", getMoreThanSongLine)
	}
}

func TestGetLyricZeroLine(t *testing.T) {
	t.Logf("Get lyric more than song line : %s", getZeroThanSongLine)
	if songTest.Recite(0) != getZeroThanSongLine {
		t.Errorf("WRONG! it must %s", getZeroThanSongLine)
	}
}

func TestGetObjectZeroLine(t *testing.T) {
	t.Logf("Get object more than song line : %s", getZeroThanSongLine)
	if songTest.ReciteObject(0) != getZeroThanSongLine {
		t.Errorf("WRONG! it must %s", getZeroThanSongLine)
	}
}
