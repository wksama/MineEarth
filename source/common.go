package source

import (
	"MineEarth/core"
	"MineEarth/system"
	"bytes"
	"github.com/nfnt/resize"
	"github.com/spf13/viper"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	perimeter uint
	lrOffset  int
	tbOffset  int
	locker    sync.WaitGroup
)

const (
	Sbing = iota
	Scustomize
	Sfengyun4
	Ssunflower8
)

var NameMap = map[int]string{
	Sbing:       "必应",
	Scustomize:  "自定义",
	Sfengyun4:   "风云4号",
	Ssunflower8: "向日葵8号",
}

var StructMap = map[int]interface{}{
	Sbing:       Bing{},
	Scustomize:  Customize{},
	Sfengyun4:   Fy{},
	Ssunflower8: Sunflower{},
}

type Base struct {
	ImageUrl     string
	ImageContent image.Image
	Canvas       *image.RGBA
	WallPaper    image.Image
}

func (b *Base) GetImageContent() error {
	resp, err := http.Get(b.ImageUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	b.ImageContent, _, err = image.Decode(bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}

	percentage := viper.GetFloat64("wallpaper.percentage")
	perimeter = uint(math.Round(float64(system.ScreenH) * percentage))
	lrOffset = (system.ScreenW - int(perimeter)) / 2
	tbOffset = (system.ScreenH - int(perimeter)) / 2

	return nil
}

func (b *Base) ResizeImage() {
	b.ImageContent = resize.Resize(perimeter, perimeter, b.ImageContent, resize.Lanczos3)
}

func (b *Base) GetCanvas() {
	//draw a new black background image which size the save as screen
	b.Canvas = image.NewRGBA(image.Rect(0, 0, system.ScreenW, system.ScreenH))
	for x := 0; x < system.ScreenW; x++ {
		for y := 0; y < system.ScreenH; y++ {
			b.Canvas.Set(x, y, color.Black)
		}
	}

	return
}

func (b *Base) GetWallpaper() image.Image {
	return b.WallPaper
}

func (b *Base) Draw() {
	offset := image.Pt(lrOffset, tbOffset)
	draw.Draw(b.Canvas, b.ImageContent.Bounds().Add(offset), b.ImageContent, image.Point{}, draw.Over)
	b.WallPaper = image.Image(b.Canvas)
}

func (b *Base) SaveToFile() (path string, err error) {
	var f *os.File
	if runtime.GOOS == "windows" {
		path = core.CacheDir + string(os.PathSeparator) + "wallpaper.bmp"
		f, err = os.Create(path)
		if err != nil {
			log.Println("创建壁纸文件失败")
			return "", err
		}
		err = bmp.Encode(f, b.WallPaper)
		if err != nil {
			log.Println("壁纸文件编码失败")
			return "", err
		}
	} else {
		path = core.CacheDir + string(os.PathSeparator) + "wallpaper.png"
		f, _ = os.Create(path)
		err = png.Encode(f, b.WallPaper)
		if err != nil {
			log.Println("壁纸文件编码失败")
			return "", err
		}
	}
	f.Close()

	return path, nil
}

func (b *Base) GetDuration() time.Duration {
	return 10 * time.Minute
}

func (b *Base) Clean() {
	b.ImageContent = nil
	b.Canvas = nil
	b.WallPaper = nil
}
