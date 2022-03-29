package types

import "github.com/pkg/errors"

type ListType int

const (
	ListTypeTop ListType = iota + 1
)

func NewListType(s string) (ListType, error) {
	switch s {
	case "top":
		return ListTypeTop, nil
	default:
		return ListType(0), errors.New("invalid list type")
	}
}

func (l ListType) String() string {
	switch l {
	case ListTypeTop:
		return "top"
	default:
		return ""
	}
}
