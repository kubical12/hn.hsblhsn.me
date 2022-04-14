package bionify

import (
	"math"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/hsblhsn/hn.hsblhsn.me/featureflags"
	"golang.org/x/exp/utf8string"
)

func moveIfMark(str *utf8string.String, index int) int {
	if charCount := str.RuneCount(); index >= charCount {
		return charCount
	}
	if unicode.IsMark(str.At(index)) {
		return moveIfMark(str, index+1)
	}
	return index
}

func word(word string) string {
	charCount := utf8.RuneCountInString(word)
	if charCount == 1 {
		return word
	}
	utf8str := utf8string.NewString(word)
	// nolint:gomnd // this is pure magic.
	numBold := int(math.Ceil(float64(charCount) * float64(0.3)))
	numBold = moveIfMark(utf8str, numBold)

	return "<b bionic-bold>" + utf8str.Slice(0, numBold) + "</b>" +
		"<span bionic-span>" + utf8str.Slice(numBold, charCount) + "</span>"
}

const MinimumTextLength = 5

func Text(text string) string {
	if !featureflags.IsOn(featureflags.FeatureBionify, false) {
		return text
	}
	charCount := utf8.RuneCountInString(text)
	if charCount < MinimumTextLength {
		return text
	}
	res := ""
	words := strings.Split(text, " ")
	for i, v := range words {
		bionified := word(v)
		if i == len(words)-1 {
			res += bionified
			continue
		}
		res += bionified + " "
	}
	return res
}
