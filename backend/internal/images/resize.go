package images

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/nfnt/resize"
)

func Resize(content io.Reader, size ImageSize) (image.Image, error) {
	img, _, err := image.Decode(content)
	if err != nil {
		return nil, err
	}
	// resize image
	width, height := size.Dimension()
	img = resize.Thumbnail(width, height, img, resize.NearestNeighbor)
	return img, nil
}
