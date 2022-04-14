package bionify

import (
	"io"
	"math"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"golang.org/x/exp/utf8string"
	"golang.org/x/net/html"
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

func Paragraph(text string) string {
	charCount := utf8.RuneCountInString(text)
	const minimumContentLength = 10
	if charCount < minimumContentLength {
		return text
	}
	res := ""
	words := strings.Split(text, " ")
	for _, word := range words {
		res += Word(word) + " "
	}
	return res
}

func Text(text string) string {
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
	output := &strings.Builder{}
	full := doc.Find("*")
	for i := range full.Nodes {
		Node(full.Nodes[i], output)
	}
	body := full.Find("body")
	body.SetHtml(output.String())
	str, err := body.Html()
	if err != nil {
		return "", errors.Wrap(err, "bionify: could not render html node to string")
	}
	return str, nil
}

func Node(node *html.Node, buf io.Writer) {
	if node == nil {
		return
	}
	if node.Type == html.TextNode {
		_, _ = buf.Write([]byte(Text(node.Data)))
		return
	}
	Node(node.FirstChild, buf)
}
