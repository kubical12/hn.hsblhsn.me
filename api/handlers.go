package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/clients"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/hackernews"
)

var (
	hn = hackernews.NewHackerNews()
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
		r.Path("/feeds/{kind:new|top}/{page:[0-9]{1,3}}").
			Methods(http.MethodGet).
			Handler(AttachMiddlewares(h.feedList, time.Second*10, time.Hour))
		r.Path("/feed_items/{id:[0-9]+}").
			Methods(http.MethodGet).
			Handler(AttachMiddlewares(h.feedItem, time.Second*3, time.Hour))
		r.Path("/feed_images").
			Queries("imageUrl", "{imageUrl}").
			Methods(http.MethodGet).
			Handler(AttachMiddlewares(h.feedImage, time.Second*3, time.Hour))
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
	feed, err := hn.GetFeed(r.Context(), feedKind, pageNum)
	if err != nil {
		HTTPError(w, err, http.StatusInternalServerError, "Failed to get feed items.")
		return
	}
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
	feedItem, err := hn.GetFeedItem(r.Context(), itemIDNum, true)
	if err != nil {
		HTTPError(w, err, http.StatusInternalServerError, "Failed to get feed item.")
		return
	}
	if err := json.NewEncoder(w).Encode(feedItem); err != nil {
		HTTPError(w, err, http.StatusInternalServerError, "Failed to encode feed item.")
		return
	}
}

func (h *Handler) feedImage(w http.ResponseWriter, r *http.Request) {
	var (
		params = mux.Vars(r)
		url    = params["imageUrl"]
	)
	reader, err := clients.SendHTTPRequest(r.Context(), url)
	if err != nil {
		HTTPError(w, err, http.StatusInternalServerError, "Failed to get image.")
		return
	}
	if _, err := io.Copy(w, reader); err != nil {
		HTTPError(w, err, http.StatusInternalServerError, "Failed to copy image.")
		return
	}
}
