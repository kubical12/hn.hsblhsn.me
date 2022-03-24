package hackernews

import "errors"

// FeedKind is an enum type for HackerNews feed kind.
type FeedKind int

const (
	FeedKindTop FeedKind = iota + 1
	FeedKindNew
)

// NewFeedKind returns a FeedKind from a string.
func NewFeedKind(str string) (FeedKind, error) {
	switch str {
	case "new":
		return FeedKindNew, nil
	case "top":
		return FeedKindTop, nil
	default:
		return FeedKind(0), errors.New("hackernews: Unknown FeedKind")
	}
}

// String returns the string representation of a FeedKind.
func (i FeedKind) String() string {
	switch i {
	case FeedKindNew:
		return "newstories"
	case FeedKindTop:
		return "topstories"
	default:
		panic(errors.New("hackernews: Unknown FeedKind"))
	}
}
