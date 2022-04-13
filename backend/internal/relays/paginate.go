package relays

import (
	"strconv"
)

type ConnectionInput struct {
	After  *string
	Before *string
	First  *int
	Last   *int
}

type Paginator struct {
	total int
	Start int
	End   int
}

func Paginate(total int, input ConnectionInput) (*Paginator, error) {
	start, end := 0, 0

	if input.First != nil && input.Last != nil {
		return nil, ErrBothFirstAndLast
	}
	if input.After != nil {
		start, _ = strconv.Atoi(*input.After)
		start = min(start+1, total)
	}

	if input.Before != nil {
		end, _ = strconv.Atoi(*input.Before)
		end = min(end, total)
	} else {
		end = total
	}

	before, after := end, start

	if start < 0 || start > end {
		return nil, ErrInvalidPaging
	}

	switch {
	case input.First != nil:
		end = start + min(*input.First, MaxCursorGap)
	case input.Last != nil:
		start = end - min(*input.Last, MaxCursorGap)
	default:
		end = start + MaxCursorGap
	}

	end = min(end, before)
	start = max(start, after)

	return &Paginator{
		total: total,
		Start: start,
		End:   end,
	}, nil
}

func (p *Paginator) HasNext() bool {
	return p.Start < p.End
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
