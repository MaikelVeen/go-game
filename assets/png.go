package assets

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"image/png"
)

func loadFile(fs embed.FS, name string) (image.Image, error) {
	content, err := staticSpritesFS.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("failed to load png: %w", err)
	}

	img, err := png.Decode(bytes.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("failed to decode png: %w", err)
	}

	return img, nil
}
