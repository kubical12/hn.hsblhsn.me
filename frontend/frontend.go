package frontend

import (
	"embed"
	"errors"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers routes for the server.
func RegisterRoutes(r *mux.Router) {
	h := &staticFileServer{
		FS: newSpaFS(assetFS, "build"),
	}
	r.PathPrefix("/").Handler(h)
}

// spaFS is an optimized filesystem implementation for single page application.
// It can serve files from any directory in an embedded fs.
// And, it resolves to /index.html if any file is not found in the fs.
type spaFS struct {
	dir      string
	internal embed.FS
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
	f, err := a.internal.Open(filepath.Join(a.dir, name))
	if errors.Is(err, fs.ErrNotExist) {
		return a.internal.Open(filepath.Join(a.dir, "index.html"))
	}
	return f, err
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
