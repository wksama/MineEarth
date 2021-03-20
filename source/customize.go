package source

import "image"

type Customize struct {
	Base
}

func (c *Customize) GetImageUrl() error {
	return nil
}

func (c *Customize) GetImageContent() error {
	return nil
}

func (c *Customize) SetWallpaper(wallpaper image.Image) {
	c.WallPaper = wallpaper
}

func (c *Customize) ResizeImage() {

}

func (c *Customize) GetCanvas() {

}

func (c *Customize) Draw() {

}
