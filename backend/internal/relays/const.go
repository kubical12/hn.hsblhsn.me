package relays

import "github.com/pkg/errors"

const (
	MaxCursorGap   = 10
	MaxConcurrency = 20
)

var (
	ErrBothFirstAndLast = errors.New("relays: cannot specify both first and last")
	ErrInvalidPaging    = errors.New("relays: invalid paging")
)
