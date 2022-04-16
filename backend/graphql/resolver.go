package graphql

import "github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/hackernews"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	hackerNews *hackernews.HackerNews
}

func NewResolver(hackerNews *hackernews.HackerNews) *Resolver {
	return &Resolver{hackerNews: hackerNews}
}
