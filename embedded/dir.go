package embedded

import (
	"embed"
	"errors"
	"io/fs"
	"path/filepath"
)

//go:embed dist/*
var assetFS embed.FS

var Assets = NewAssetFS(assetFS, "dist")

type assets struct {
	dir      string
	internal embed.FS
}

func NewAssetFS(root fs.FS, dir string) fs.FS {
	return &assets{
		dir:      dir,
		internal: assetFS,
	}
}

func (a *assets) Open(name string) (fs.File, error) {
	f, err := a.internal.Open(filepath.Join(a.dir, name))
	if errors.Is(err, fs.ErrNotExist) {
		return a.internal.Open(filepath.Join(a.dir, "index.html"))
	}
	return f, err
}
