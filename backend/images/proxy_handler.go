package images

import (
	"image"
	"image/jpeg"
	"net/http"

	"github.com/tasylab/hn.hsblhsn.me/backend/internal/httpclient"
	"github.com/tasylab/hn.hsblhsn.me/backend/internal/logutil"
	"go.uber.org/zap"
)

const quality = 85

type ImageProxyHandler struct {
	client *httpclient.Client
	logger *zap.Logger
}

func NewImageProxyHandler(client *httpclient.Client, logger *zap.Logger) *ImageProxyHandler {
	return &ImageProxyHandler{
		client: client,
		logger: logger.With(zap.String("component", "images_resize_handler")),
	}
}

func (h *ImageProxyHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// get image size from query string
	size := ImageSize(logutil.Sanitize(req.URL.Query().Get("size")))
	src := logutil.Sanitize(req.URL.Query().Get("src"))
	if size == "" {
		size = ImageSizeFull
	}
	logger := h.logger.With(zap.Stringer("size", size), zap.String("src", src))

	// get image src from query string
	if src == "" {
		logger.Error("no src provided")
		writeBlankImage(resp, logger)
		return
	}
	// get image from src
	imgResp, err := h.client.Get(req.Context(), src)
	if err != nil {
		logger.Error("failed to get image", zap.Error(err))
		writeBlankImage(resp, logger)
		return
	}
	// resize image
	resized, err := Resize(imgResp.Body, size)
	if err != nil {
		logger.Error("failed to resize image", zap.Error(err))
		writeBlankImage(resp, logger)
		return
	}

	err = jpeg.Encode(resp, resized, &jpeg.Options{
		Quality: quality,
	})
	if err != nil {
		logger.Error("failed to encode image", zap.Error(err))
		writeBlankImage(resp, logger)
		return
	}
}

func blankImage() image.Image {
	return image.NewRGBA(image.Rect(0, 0, 1, 1))
}

func writeBlankImage(w http.ResponseWriter, logger *zap.Logger) {
	w.WriteHeader(http.StatusOK)
	if err := jpeg.Encode(w, blankImage(), &jpeg.Options{Quality: quality}); err != nil {
		logger.Error("failed to write blank image", zap.Error(err))
	}
}
