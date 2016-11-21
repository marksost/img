package mutableimages

type (
	StaticMutableImage struct {
		img *ProcessableImage
	}
)

func NewStaticMutableImage(img *ProcessableImage) *StaticMutableImage {
	return &StaticMutableImage{
		img: img,
	}
}

func (i *StaticMutableImage) GetWidth() int64 {
	return 0
}

func (i *StaticMutableImage) GetHeight() int64 {
	return 0
}

func (i *StaticMutableImage) SetDefaults() {}
