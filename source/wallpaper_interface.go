package source

import (
	"image"
	"time"
)

type WallpaperInterface interface {
	GetImageUrl() error
	GetImageContent() error
	ResizeImage()
	GetCanvas()
	Draw()
	GetWallpaper() image.Image
	Clean()
	SaveToFile() (path string, err error)
	GetDuration() time.Duration
}
