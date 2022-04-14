package frontend

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// RegisterRoutes registers routes for the server.
func RegisterRoutes(r *mux.Router) {
	h := &staticFileServer{
		FS: newSpaFS(assetFS, "build"),
	}
	r.PathPrefix("/").Handler(prerender(h))
}

// spaFS is an optimized filesystem implementation for single page application.
// It can serve files from any directory in an embedded fs.
// And, it resolves to /index.html if any file is not found in the fs.
type spaFS struct {
	internal embed.FS
	dir      string
}

// newSpaFS returns a new fs that roots to th given dir.
func newSpaFS(root fs.FS, dir string) fs.FS {
	return &spaFS{
		dir:      dir,
		internal: assetFS,
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

func prerender(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		ua := strings.ToLower(request.UserAgent())
		isBot := strings.Contains(ua, "bot") && !strings.Contains(ua, "google")
		if !isBot {
			next.ServeHTTP(response, request)
			return
		}

		request.URL.Host = request.Host
		endpoint := fmt.Sprintf("https://service.prerender.cloud/%s", url.PathEscape(request.URL.String()))
		log.Println("frontend: prerendering endpoint", endpoint)
		req, err := http.NewRequestWithContext(request.Context(), http.MethodGet, endpoint, nil)
		if err != nil {
			log.Println("frontend: could not create request", err)
			next.ServeHTTP(response, request)
			return
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("frontend: could not send request", err)
			next.ServeHTTP(response, request)
			return
		}
		defer resp.Body.Close()
		response.WriteHeader(resp.StatusCode)
		_, _ = io.Copy(response, resp.Body)
	})
}
