package src

import (
	"fmt"
	"math/rand"
	"strings"
)

// Object
type Song struct {
	Lyric []string // save the lyric per line using array
}

// Funcrion to get one lyric line by index
func (s Song) getLine(index int) string {
	return s.Lyric[index]
}

// Function to get sum of the lyric line
func (s Song) getSize() int {
	return len(s.Lyric)
}

// Function to get the lyric from n-last line
func (s Song) getLyric(index int) (res string) {
	for i := s.getSize() - index; i < s.getSize(); i++ { // looping all line from n-last line until last line
		res += " " + s.getLine(i) // combine the lyric line
	}
	return res
}

// Function to get the object from n-last line
func (s Song) getObject(index int) (res string) {
	for i := s.getSize() - index; i < s.getSize(); i++ { // looping all line from n-last line until last line
		splitString := strings.Split(s.getLine(i), " ") // split string by whitespace
		if index == 1 { // special case to handle when we want to get the last object of the lyric line
			res += "the " + splitString[1] // combine the object of the lyric line
		} else {
			if i != s.getSize()-1 { // this condition to handle when we must add ',' or 'and'
				res += "the " + splitString[1] + ", " // combine the object of the lyric line
			} else {
				res = res[:len(res)-2]
				res += " and the " + splitString[1] // combine the object of the lyric line
			}
		}
	}
	return res
}

// Function to recite the song
func (s Song) Recite(input ...int) string {
	if input == nil { // check if input is empty
		Randomizer := rand.Intn(s.getSize()-1) + 1 // create random number from 1 until sum of the lyric line -1
		res := s.getLyric(Randomizer) // call getLyric function using random number
		return res[1:]
	}

	if input[0] > s.getSize() { // validating input cannot more than sum of the lyric
		return fmt.Sprintf("Input more than %d song line", s.getSize())
	}
	if input[0] == 0 { // validating input cannot zero
		return "Input cannot zero"
	}

	res := s.getLyric(input[0]) // call getLyric function using input number
	return res[1:]
}

// Function to recite the song (object only)
func (s Song) ReciteObject(input ...int) string {
	if input == nil { // check if input is empty
		randomizer := rand.Intn(s.getSize()-1) + 1 // create random number from 1 until sum of the lyric line -1
		res := s.getObject(randomizer) // call getObject function using random number
		return "This is " + res
	}

	if input[0] > s.getSize() { // validating input cannot more than sum of the lyric
		return fmt.Sprintf("Input more than %d song line", s.getSize())
	}
	if input[0] == 0 { // validating input cannot zero
		return "Input cannot zero"
	}

	res := s.getObject(input[0]) // call getObject function using input number
	return "This is " + res
}