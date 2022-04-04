package bionify

import (
	"math"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/exp/utf8string"
)

func Word(word string) string {
	charCount := utf8.RuneCountInString(word)
	if charCount == 1 {
		return word
	}
	utf8str := utf8string.NewString(word)
	numBold := int(math.Ceil(float64(charCount) * float64(0.3)))
	numBold = moveIfMark(utf8str, numBold)

	return "<b bionic-bold>" + utf8str.Slice(0, numBold) + "</b>" +
		"<span bionic-span>" + utf8str.Slice(numBold, charCount) + "</span>"
}

func moveIfMark(str *utf8string.String, index int) int {
	charCount := str.RuneCount()
	if index >= charCount {
		return charCount
	}
	if unicode.IsMark(str.At(index)) {
		return moveIfMark(str, index+1)
	}
	return index
}

func Text(text string) string {
	charCount := utf8.RuneCountInString(text)
	if charCount < 10 {
		return text
	}
	res := ""
	words := strings.Split(text, " ")
	for _, word := range words {
		res += Word(word) + " "
	}
	return res
}
