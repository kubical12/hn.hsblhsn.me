package frontend

import (
	"embed"
)

//go:embed build/*
var assetFS embed.FS
