package types

type List struct {
	Type  ListType `json:"type"`
	Page  uint8    `json:"page"`
	Items []*Item  `json:"items"`
	SEO   *SEO     `json:"seo"`
}
