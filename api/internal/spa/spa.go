package spa

import (
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

type Asset struct {
	fs.FS
}

func RegisterRoutes(r *mux.Router, source fs.FS) {
	h := &Asset{
		FS: source,
	}
	r.PathPrefix("/").Handler(h)
}

func (a *Asset) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.FS(a)).ServeHTTP(w, r)
}
