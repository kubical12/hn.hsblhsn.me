package images

import (
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/logutil"
	"image/jpeg"
	"net/http"

	"go.uber.org/zap"
)

type SocialPreviewHandler struct {
	generator *SocialPreviewGenerator
	logger    *zap.Logger
}

func NewSocialPreviewHandler(generator *SocialPreviewGenerator, logger *zap.Logger) *SocialPreviewHandler {
	return &SocialPreviewHandler{
		generator: generator,
		logger:    logger,
	}
}

func (h *SocialPreviewHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	title := req.URL.Query().Get("title")
	if title == "" {
		title = "Hackernews client focused on content and readability"
	}
	logger := h.logger.With(
		zap.String("component", "social_preview"),
		zap.String("title", logutil.Sanitize(title)),
	)
	img, err := h.generator.Generate(title)
	if err != nil {
		logger.Error(
			"failed to encode generated social preview image",
			zap.Error(err),
		)
		writeBlankImage(resp, logger)
		return
	}
	err = jpeg.Encode(resp, img, &jpeg.Options{
		Quality: 85,
	})
	if err != nil {
		logger.Error(
			"failed to encode generated social preview image",
			zap.Error(err),
		)
		writeBlankImage(resp, logger)
		return
	}
}
