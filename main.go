package main

import (
	"fmt"
	entity "recite/src"
)

func main() {
	aSong := entity.Song{Lyric: []string{
		"the horse and the hound and the horn that belonged to",
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
	
	fmt.Println(aSong.Recite(2))
	fmt.Println(aSong.ReciteObject(1))
}
