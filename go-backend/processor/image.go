package processor

import (
	"context"
	"image"
	"os"

	"github.com/fogleman/gg"
)

type ImageProcessor interface {
	AddTextToImage(ctx context.Context, originalpath, resultpath, text string) error
}

type imageProcessor struct {
}

// AddTextToImage implements ImageProcessor.
func (i *imageProcessor) AddTextToImage(ctx context.Context, originalpath string, resultpath string, text string) error {
	f, err := os.Open(originalpath)
	if err != nil {
		return err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	dc := gg.NewContextForImage(img)
	if err = dc.LoadFontFace("/Users/kennethjohn/Personal Docs/projects/caption-canvas/go-backend/fonts/Nunito_Sans/static/NunitoSans_7pt_SemiCondensed-ExtraBold.ttf", 60); err != nil {
		return err
	}

	dc.SetRGB(1, 1, 1)
	const W = 1024
	const H = 1024
	const P = 16
	// dc.DrawStringWrapped(text, 512, 900, 1, 0.5, 500, 0.5, gg.AlignCenter)
	dc.DrawStringAnchored(text, 512, 916, 0.5, 0.5)

	err = dc.SavePNG(resultpath)
	if err != nil {
		return err
	}

	return nil
}

func NewImageProcessor() ImageProcessor {
	return &imageProcessor{}
}
