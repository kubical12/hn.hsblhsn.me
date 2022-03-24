package hackernews

// Feed is a collection of posts.
type Feed struct {
	FeedItems []*FeedItem `json:"feedItems"`
}
