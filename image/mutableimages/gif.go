package mutableimages

type (
	GifMutableImage struct {
		img *ProcessableImage
	}
)

func NewGifMutableImage(img *ProcessableImage) (*GifMutableImage, error) {
	return &GifMutableImage{
		img: img,
	}, nil
}

func (i *GifMutableImage) GetWidth() int64 {
	return 0
}

func (i *GifMutableImage) GetHeight() int64 {
	return 0
}

func (i *GifMutableImage) SetDefaults() {}
