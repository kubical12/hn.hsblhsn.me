package types

import (
	"context"
	"encoding/json"
	"image"
	"image/jpeg"
	"log"
	"net/http"
)

type GetListByTypeResponse = List

func EncodeGetListByTypeResponse(ctx context.Context, w http.ResponseWriter, d any) error {
	return encodeJSON(ctx, w, d)
}

type GetItemByIDResponse = Item

func EncodeGetItemByIDResponse(ctx context.Context, w http.ResponseWriter, d any) error {
	return encodeJSON(ctx, w, d)
}

type GetResizedImageResponse = image.Image

func EncodeGetResizedImageResponse(_ context.Context, w http.ResponseWriter, d any) error {
	response, ok := d.(image.Image)
	if !ok {
		log.Printf("unexpected type in image encoder %T", d)
	}
	return jpeg.Encode(w, response, &jpeg.Options{
		Quality: 90,
	})
}

func encodeJSON(_ context.Context, w http.ResponseWriter, response any) error {
	return json.NewEncoder(w).Encode(response)
}
