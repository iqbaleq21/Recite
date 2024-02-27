package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var Lyric = []string{
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
}
var DataSize = len(Lyric)

func recite(input ...int) {
	Output := ""

	if input == nil {
		Randomizer := rand.Intn(DataSize)
		for i := DataSize - Randomizer; i < DataSize; i++ {
			Output += " " + Lyric[i]
		}
		fmt.Println(Output[1:])
		return
	}

	if input[0] > 12 {
		fmt.Println("Input lebih besar dari 12! Coba lagi!")
		return
	}

	for i := DataSize - input[0]; i < DataSize; i++ {
		Output += " " + Lyric[i]
	}

	fmt.Println(Output[1:])
}

func reciteObject(input ...int) {
	Output := "This is "

	if input == nil {
		Randomizer := rand.Intn(DataSize)
		for i := DataSize - Randomizer; i < DataSize; i++ {
			splitString := strings.Split(Lyric[i], " ")
			if Randomizer == 1 {
				Output += "the " + splitString[1]
			} else {
				if i != DataSize-1 {
					Output += "the " + splitString[1] + ", "
				} else {
					Output = Output[:len(Output)-2]
					Output += " and the " + splitString[1]
				}
			}
		}
		fmt.Println(Output)
		return
	}

	if input[0] > 12 {
		fmt.Println("Input lebih besar dari 12! Coba lagi!")
		return
	}

	for i := DataSize - input[0]; i < DataSize; i++ {
		splitString := strings.Split(Lyric[i], " ")
		if input[0] == 1 {
			Output += "the " + splitString[1]
		} else {
			if i != DataSize-1 {
				Output += "the " + splitString[1] + ", "
			} else {
				Output = Output[:len(Output)-2]
				Output += " and the " + splitString[1]
			}
		}
	}

	fmt.Println(Output)
}

func main() {
	recite(2)
	recite(99)
	recite()

	reciteObject(7)
	reciteObject(21)
	reciteObject()
}
