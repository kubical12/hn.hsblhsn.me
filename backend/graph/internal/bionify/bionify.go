package bionify

import (
	"math"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"golang.org/x/exp/utf8string"
)

func Word(word string) string {
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

func moveIfMark(str *utf8string.String, index int) int {
	if charCount := str.RuneCount(); index >= charCount {
		return charCount
	}
	if unicode.IsMark(str.At(index)) {
		return moveIfMark(str, index+1)
	}
	return index
}

func Text(text string) string {
	charCount := utf8.RuneCountInString(text)
	const minimumSentenceLength = 10
	if charCount < minimumSentenceLength {
		return text
	}
	res := ""
	words := strings.Split(text, " ")
	for _, word := range words {
		res += Word(word) + " "
	}
	return res
}

func HTMLText(content string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return "", errors.Wrap(err, "bionify: could not parse html from reader")
	}
	all := doc.Find("*").Each(func(i int, s *goquery.Selection) {
		children := s.Children()
		if len(children.Nodes) == 0 {
			s.SetHtml(Text(s.Text()))
		}
	})
	body := all.Find("body")
	str, err := body.Html()
	if err != nil {
		return "", errors.Wrap(err, "bionify: could not render html node to string")
	}
	return str, nil
}
