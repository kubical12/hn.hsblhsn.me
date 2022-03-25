package api

import (
	"context"
	"encoding/json"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/elnormous/contenttype"
	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/clients"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/hackernews"
)

type Handler struct {
	hnClient *hackernews.HackerNews
	once     sync.Once
}

func NewHandler(client *hackernews.HackerNews) *Handler {
	return &Handler{
		hnClient: client,
		once:     sync.Once{},
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	h.once.Do(func() {
		const (
			DefaultReqTimeout  = time.Second * 10
			DefaultCacheMaxAge = time.Minute * 30
		)
		r.Path("/feeds/{kind:new|top}/{page:[0-9]{1,3}}").
			Methods(http.MethodGet).
			Handler(AttachMiddlewares(h.feedList, DefaultReqTimeout, DefaultCacheMaxAge))
		r.Path("/feed_items/{id:[0-9]+}").
			Methods(http.MethodGet).
			Handler(AttachMiddlewares(h.feedItem, DefaultReqTimeout, DefaultCacheMaxAge))
		r.Path("/feed_images").
			Queries("imageUrl", "{imageUrl}").
			Methods(http.MethodGet).
			Handler(AttachMiddlewares(h.feedImage, DefaultReqTimeout, DefaultCacheMaxAge))
	})
}

func (h *Handler) feedList(w http.ResponseWriter, r *http.Request) {
	var (
		params = mux.Vars(r)
		kind   = params["kind"]
		page   = params["page"]
	)
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		HTTPError(w, err, http.StatusBadRequest, "Invalid page number.")
		return
	}
	feedKind, err := hackernews.NewFeedKind(kind)
	if err != nil {
		HTTPError(w, err, http.StatusBadRequest, "Invalid feed kind.")
		return
	}
	feed, err := h.hnClient.GetFeed(r.Context(), feedKind, pageNum)
	if err != nil {
		HTTPError(w, err, http.StatusInternalServerError, "Failed to get feed items.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(feed); err != nil {
		HTTPError(w, err, http.StatusInternalServerError, "Failed to encode feed items.")
		return
	}
}

func (h *Handler) feedItem(w http.ResponseWriter, r *http.Request) {
	var (
		params = mux.Vars(r)
		itemID = params["id"]
	)
	itemIDNum, err := strconv.Atoi(itemID)
	if err != nil {
		HTTPError(w, err, http.StatusBadRequest, "Invalid item ID.")
		return
	}
	feedItem, err := h.hnClient.GetFeedItem(r.Context(), itemIDNum, true)
	if err != nil {
		HTTPError(w, err, http.StatusInternalServerError, "Failed to get feed item.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(feedItem); err != nil {
		HTTPError(w, err, http.StatusInternalServerError, "Failed to encode feed item.")
		return
	}
}

func (h *Handler) feedImage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()
	r = r.WithContext(ctx)

	var (
		params   = mux.Vars(r)
		url      = params["imageUrl"]
		fallback = params["fallback"]
	)
	if fallback == "true" {
		ImgErr(w, nil, http.StatusOK, "Fallback image requested")
		return
	}

	reader, err := clients.SendHTTPRequest(r.Context(), url)
	if err != nil {
		ImgErr(w, err, http.StatusInternalServerError, "Failed to get image.")
		return
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		ImgErr(w, err, http.StatusInternalServerError, "Failed to decode image.")
		return
	}

	availableMediaTypes := []contenttype.MediaType{
		contenttype.NewMediaType("image/jpeg"),
		contenttype.NewMediaType("image/png"),
		contenttype.NewMediaType("image/apng"),
	}
	accepted, _, err := contenttype.GetAcceptableMediaType(r, availableMediaTypes)
	if err != nil {
		ImgErr(w, err, http.StatusInternalServerError, "Please pass a valid Accept header")
		return
	}

	if accepted.Subtype == "png" {
		w.Header().Set("Content-Type", "image/png")
		err = png.Encode(w, img)
	} else if accepted.Type == "image" || accepted.Subtype == "jpeg" {
		w.Header().Set("Content-Type", "image/jpeg")
		err = jpeg.Encode(w, img, nil)
	}

	if err != nil {
		ImgErr(w, err, http.StatusInternalServerError, "Failed to encode image.")
		return
	}
}
