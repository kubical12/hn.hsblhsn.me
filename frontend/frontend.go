package frontend

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// RegisterRoutes registers routes for the server.
func RegisterRoutes(r *mux.Router, logger *zap.Logger) {
	h := &staticFileServer{
		FS: newSpaFS(assetFS, "build"),
	}
	r.PathPrefix("/").
		Handler(PrerenderIfBot(h, logger))
}

// spaFS is an optimized filesystem implementation for single page application.
// It can serve files from any directory in an embedded fs.
// And, it resolves to /index.html if any file is not found in the fs.
type spaFS struct {
	internal embed.FS
	dir      string
}

// newSpaFS returns a new fs that roots to th given dir.
func newSpaFS(root embed.FS, dir string) fs.FS {
	return &spaFS{
		dir:      dir,
		internal: root,
	}
}

// Open implements fs.FS.
func (a *spaFS) Open(name string) (fs.File, error) {
	file, err := a.internal.Open(filepath.Join(a.dir, name))
	if errors.Is(err, fs.ErrNotExist) {
		// nolint:wrapcheck // too much nested code.
		return a.internal.Open(filepath.Join(a.dir, "index.html"))
	}
	if err != nil {
		return nil, errors.Wrap(err, "frontend: could not open file")
	}
	return file, nil
}

// staticFileServer is an implementation of http.Handler,
// to serve file from a fs.FS.
type staticFileServer struct {
	fs.FS
}

// ServeHTTP implements http.Handler.
func (f *staticFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.FS(f)).ServeHTTP(w, r)
}

func PrerenderIfBot(next http.Handler, logger *zap.Logger) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// pass all non-bot requests.
		if !isBot(request.UserAgent()) {
			next.ServeHTTP(response, request)
			return
		}

		// process bot requests here.
		request.URL.Host = request.Host
		request.URL.Scheme = "https"
		endpoint := fmt.Sprintf("https://service.prerender.cloud/%s", request.URL.String())
		logger.Debug("frontend: prerendering endpoint", zap.String("url", endpoint))

		req, err := http.NewRequestWithContext(request.Context(), http.MethodGet, endpoint, nil)
		if err != nil {
			logger.Error("frontend: could not create request", zap.Error(err))
			next.ServeHTTP(response, request)
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			logger.Error("frontend: could not prerender", zap.Error(err))
			next.ServeHTTP(response, request)
			return
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
		if resp.StatusCode != http.StatusOK {
			logger.Error("frontend: prerender failed", zap.Int("status", resp.StatusCode))
			next.ServeHTTP(response, request)
			return
		}
		response.WriteHeader(resp.StatusCode)
		_, _ = io.Copy(response, resp.Body)
	})
}

func isBot(useragent string) bool {
	useragent = strings.ToLower(useragent)
	// google bot can render react page.
	// no need to prerender.
	if strings.Contains(useragent, "+http://www.google.com/bot.html") {
		return false
	}
	bots := []string{
		"bot",
		"facebookexternalhit",
		"embedly",
		"wordpress",
		"curl",
		"Go-http-client",
	}
	for _, bot := range bots {
		if strings.Contains(useragent, bot) {
			return true
		}
	}
	return false
}
