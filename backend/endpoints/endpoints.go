package endpoints

import (
	"context"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/services"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/types"
)

type Endpoint interface {
	GetListByType(ctx context.Context, request any) (any, error)
	GetItemByID(ctx context.Context, request any) (any, error)
	GetResizedImage(ctx context.Context, request any) (any, error)
}

type endpoint struct {
	svc services.Service
}

var _ Endpoint = (*endpoint)(nil)

func NewEndpoints(svc services.Service) Endpoint {
	return &endpoint{svc: svc}
}

func (h *endpoint) GetListByType(ctx context.Context, request any) (any, error) {
	req := request.(*types.GetListByTypeRequest)
	itemList, err := h.svc.GetListByType(ctx, req.Type, req.Page)
	if err != nil {
		return nil, err
	}
	return itemList, nil
}

func (h *endpoint) GetItemByID(ctx context.Context, request any) (any, error) {
	req := request.(*types.GetItemByIDRequest)
	item, err := h.svc.GetItemByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (h *endpoint) GetResizedImage(ctx context.Context, request any) (any, error) {
	req := request.(*types.GetResizedImage)
	img, err := h.svc.GetResizedImage(ctx, req.URL, req.Size)
	if err != nil {
		return nil, err
	}
	return img, nil
}
