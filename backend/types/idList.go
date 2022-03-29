package types

const (
	ItemsPerPageLimit int = 10
)

// IDList is a list of IDs.
// This is a wrapper around a slice of ints.
// It provides pagination support over a slice.
type IDList []int32

// NewIDList returns a new IDList from the given slice.
func NewIDList(list []int32) IDList {
	return IDList(list)
}

// Len returns the length of the list.
func (l IDList) Len() int {
	return len(l)
}

// Paginate returns a slice of IDs from the list.
// It starts at the given page and returns the given number of IDs.
func (l IDList) Paginate(page int) []int32 {
	if page < 1 {
		return []int32{}
	}
	var (
		limit  = ItemsPerPageLimit
		offset = (page - 1) * limit
	)
	var (
		start = offset
		end   = offset + limit
	)
	if start >= l.Len() {
		return []int32{}
	}
	if end > l.Len() {
		end = l.Len()
	}
	return l[start:end]
}
