package types

type Item struct {
	ID            uint32 `json:"id" mapstructure:"id,omitempty"`
	Title         string `json:"title" mapstructure:"title"`
	Summary       string `json:"summary"`
	Content       string `json:"content"`
	ThumbnailUrl  string `json:"thumbnailUrl"`
	HackerNewsUrl string `json:"hackerNewsUrl"`
	Domain        string `json:"domain"`
	URL           string `json:"url" mapstructure:"url"`
	TotalComments uint16 `json:"totalComments" mapstructure:"descendants"`
	SEO           *SEO   `json:"seo"`
}
