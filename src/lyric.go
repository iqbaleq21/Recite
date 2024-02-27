package src

import (
	"fmt"
	"math/rand"
	"strings"
)

type Song struct {
	Lyric []string
}

func (s Song) getLine(index int) string {
	return s.Lyric[index]
}

func (s Song) getSize() int {
	return len(s.Lyric)
}

func (s Song) getLyric(index int) (res string) {
	for i := s.getSize() - index; i < s.getSize(); i++ {
		res += " " + s.getLine(i)
	}
	return res
}

func (s Song) getObject(index int) (res string) {
	for i := s.getSize() - index; i < s.getSize(); i++ {
		splitString := strings.Split(s.getLine(i), " ")
		if index == 1 {
			res += "the " + splitString[1]
		} else {
			if i != s.getSize()-1 {
				res += "the " + splitString[1] + ", "
			} else {
				res = res[:len(res)-2]
				res += " and the " + splitString[1]
			}
		}
	}
	return res
}

func (s Song) Recite(input ...int) string {
	if input == nil {
		Randomizer := rand.Intn(s.getSize())
		res := s.getLyric(Randomizer)
		return res[1:]
	}

	if input[0] > s.getSize() {
		return fmt.Sprintf("Input more than %d song line", s.getSize())
	}

	res := s.getLyric(input[0])
	return res[1:]
}

func (s Song) ReciteObject(input ...int) string {
	if input == nil {
		randomizer := rand.Intn(s.getSize())
		res := s.getObject(randomizer)
		return res
	}

	if input[0] > s.getSize() {
		return fmt.Sprintf("Input more than %d song line", s.getSize())
	}

	res := s.getObject(input[0])
	return "This is " + res
}