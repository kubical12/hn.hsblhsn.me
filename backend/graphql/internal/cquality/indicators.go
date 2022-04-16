package cquality

import "strings"

type Indicators []string

func NewIndicators(indicators []string) Indicators {
	return indicators
}

func (list Indicators) Indicates(text []byte) bool {
	content := string(text)
	for _, v := range list {
		if strings.Contains(content, v) {
			return true
		}
	}
	return false
}
