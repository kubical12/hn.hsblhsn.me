package relays

import (
	"strconv"
)

type Connection[T any] struct {
	PageInfo   *PageInfo  `json:"pageInfo"`
	Edges      []*Edge[T] `json:"edges"`
	TotalCount int        `json:"totalCount"`
}

type PageInfo struct {
	PageCursor      string `json:"pageCursor"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
}

type Edge[T any] struct {
	Node   T      `json:"node"`
	Cursor string `json:"cursor"`
}

func NewCursor(index int) string {
	return strconv.Itoa(index)
}

func NewID(id int) string {
	return strconv.Itoa(id)
}
