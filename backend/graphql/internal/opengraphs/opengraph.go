package opengraphs

import (
	"bytes"
	"strings"

	"github.com/otiai10/opengraph/v2"
	"github.com/pkg/errors"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/internal/readerviews"
	"github.com/tasylab/hn.hsblhsn.me/backend/images"
)

type (
	OpenGraph = opengraph.OpenGraph
	Image     = opengraph.Image
	Favicon   = opengraph.Favicon
)

//nolint:gochecknoglobals // global client to call from ExternalContentLoader.
var replacer = strings.NewReplacer(
	"&nbsp;", " ",
	"&amp;", "&",
	"&quot;", "\"",
	"&#39;", "'",
)

func GetOpengraphData(uri string, content *bytes.Buffer) (*OpenGraph, error) {
	data := opengraph.New(uri)
	if err := data.Parse(content); err != nil {
		return nil, errors.Wrap(err, "opengraph: could not parse opengraph data")
	}
	if err := data.ToAbs(); err != nil {
		return nil, errors.Wrap(err, "opengraph: could not convert uri to absolute")
	}
	var err error
	data.Description, err = readerviews.TransformHTML(uri, strings.NewReader(data.Description))
	data.Description = replacer.Replace(data.Description)
	if err != nil {
		return nil, errors.Wrap(err, "opengraph: could not transform description")
	}
	for i := range data.Image {
		data.Image[i].URL = images.ProxiedURL(data.Image[i].URL, images.ImageSizeThumbnail)
	}
	if data.URL == "" {
		data.URL = uri
	}

	return data, nil
}
