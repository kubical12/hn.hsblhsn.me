package types

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/images"
	"github.com/pkg/errors"
)

type GetListByTypeRequest struct {
	Type ListType
	Page uint8
}

func DecodeGetListByTypeRequest(_ context.Context, req *http.Request) (any, error) {
	params := mux.Vars(req)
	typeParam := params["type"]
	if typeParam == "" {
		return nil, NewEncodeableError(
			errors.Wrap(ErrInvalidRequest, "types: missing list type"),
			"Invalid request parameter. missing list type.",
		)
	}
	typ, err := NewListType(typeParam)
	if err != nil {
		return nil, NewEncodeableError(
			errors.Wrap(ErrInvalidRequest, "types: invalid list type"),
			"Invalid request parameter. invalid list type.",
		)
	}
	pageParam := params["page"]
	if pageParam == "" {
		pageParam = "1"
	}
	page, err := strconv.ParseUint(pageParam, 10, 8)
	if err != nil {
		return nil, NewEncodeableError(
			errors.Wrap(ErrInvalidRequest, "types: page is not int8"),
			"Invalid request parameter. invalid page number.",
		)
	}
	return &GetListByTypeRequest{
		Type: typ,
		Page: uint8(page),
	}, nil
}

type GetItemByIDRequest struct {
	ID uint32
}

func DecodeGetItemByIDRequest(_ context.Context, req *http.Request) (any, error) {
	idParam := mux.Vars(req)["id"]
	if idParam == "" {
		return nil, NewEncodeableError(
			errors.Wrap(ErrInvalidRequest, "types: missing item id"),
			"Invalid request parameter. missing item id.",
		)
	}
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return nil, err
	}
	return &GetItemByIDRequest{
		ID: uint32(id),
	}, nil
}

type GetResizedImage struct {
	URL  string
	Size images.ImageSize
}

func DecodeGetResizedImage(_ context.Context, req *http.Request) (any, error) {
	imageURLParam := mux.Vars(req)["src"]
	if imageURLParam == "" {
		return nil, NewEncodeableError(
			errors.Wrap(ErrInvalidRequest, "types: missing image src"),
			"Invalid request parameter. missing image src",
		)
	}
	sizeParam := mux.Vars(req)["size"]
	if sizeParam == "" {
		sizeParam = images.ImageSizeFull.String()
	}
	return &GetResizedImage{
		URL:  imageURLParam,
		Size: images.ImageSize(sizeParam),
	}, nil
}
