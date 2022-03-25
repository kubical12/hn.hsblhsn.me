package api

import (
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
	"github.com/nfnt/resize"
)

// Handler for HackerNews API.
type Handler struct {
	hnClient *hackernews.HackerNews
	once     sync.Once
}

// NewHandler returns a new handler for HackerNews API.
func NewHandler(client *hackernews.HackerNews) *Handler {
	return &Handler{
		hnClient: client,
		once:     sync.Once{},
	}
}

// RegisterRoutes registers API routes.
func (h *Handler) RegisterRoutes(r *mux.Router) {
	h.once.Do(func() {
		r.Path("/feeds/{kind:new|top}/{page:[0-9]{1,3}}").
			Methods(http.MethodGet).
			Handler(cacheResp(h.feedList, DefaultCacheMaxAge))
		r.Path("/feed_items/{id:[0-9]+}").
			Methods(http.MethodGet).
			Handler(cacheResp(h.feedItem, DefaultCacheMaxAge))
		r.Path("/feed_images").
			Queries("imageUrl", "{imageUrl}", "size", "{size:full|thumbnail}").
			Methods(http.MethodGet).
			Handler(cacheResp(h.feedImage, time.Hour*72))
	})
}

// feedList serves a list of items for feeds.
func (h *Handler) feedList(w http.ResponseWriter, r *http.Request) {
	r, cancel := setRequestTimeout(r, time.Second*5)
	defer cancel()

	// get request parameters
	var (
		params = mux.Vars(r)
		kind   = params["kind"]
		page   = params["page"]
	)

	// parse page number and feed kind.
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

	// get items
	feed, err := h.hnClient.GetFeed(r.Context(), feedKind, pageNum)
	if err != nil {
		HTTPError(w, err, http.StatusInternalServerError, "Failed to get feed items.")
		return
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(feed); err != nil {
		HTTPError(w, err, http.StatusInternalServerError, "Failed to encode feed items.")
		return
	}
}

// feedItem serves a single item for feeds.
func (h *Handler) feedItem(w http.ResponseWriter, r *http.Request) {
	r, cancel := setRequestTimeout(r, time.Second*5)
	defer cancel()

	// get request parameters
	var (
		params = mux.Vars(r)
		itemID = params["id"]
	)

	// parse item id
	itemIDNum, err := strconv.Atoi(itemID)
	if err != nil {
		HTTPError(w, err, http.StatusBadRequest, "Invalid item ID.")
		return
	}

	// get item
	feedItem, err := h.hnClient.GetFeedItem(r.Context(), itemIDNum, true)
	if err != nil {
		HTTPError(w, err, http.StatusInternalServerError, "Failed to get feed item.")
		return
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(feedItem); err != nil {
		HTTPError(w, err, http.StatusInternalServerError, "Failed to encode feed item.")
		return
	}
}

// feedImage serves images for feeds.
// It acts like a proxy that resizes images that requested.
func (h *Handler) feedImage(w http.ResponseWriter, r *http.Request) {
	r, cancel := setRequestTimeout(r, time.Second*5)
	defer cancel()
	// get request parameters
	var (
		params = mux.Vars(r)
		url    = params["imageUrl"]
		size   = hackernews.ImageSize(params["size"])
	)

	// get image from url
	reader, err := clients.SendHTTPRequest(r.Context(), url)
	if err != nil {
		ImgErr(w, err, http.StatusInternalServerError, "Failed to get image.")
		return
	}

	// decode image
	img, _, err := image.Decode(reader)
	if err != nil {
		ImgErr(w, err, http.StatusInternalServerError, "Failed to decode image.")
		return
	}
	// resize image
	width, height := size.Dimension()
	img = resize.Thumbnail(width, height, img, resize.NearestNeighbor)

	// get accepted image formats
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
	// encode image based on accepted format
	if accepted.Subtype == "png" {
		w.Header().Set("Content-Type", "image/png")
		err = png.Encode(w, img)
	} else if accepted.Type == "image" || accepted.Subtype == "jpeg" {
		w.Header().Set("Content-Type", "image/jpeg")
		err = jpeg.Encode(w, img, nil)
	}
	// return error if encoding failed
	if err != nil {
		ImgErr(w, err, http.StatusInternalServerError, "Failed to encode image.")
		return
	}
}
