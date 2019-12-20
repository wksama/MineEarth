package main

import (
	"bytes"
	"earth/system"
	"earth/utils"
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"runtime"
	"time"
)

type Wallpaper struct {
	Width         int
	Height        int
	Content       []byte
	ContentLength int
}

func NewWallpaper() *Wallpaper {

	wallpaper := new(Wallpaper)
	wallpaper.Height = system.ScreenH
	wallpaper.Width = system.ScreenW
	wallpaper.ContentLength = int(math.Round(float64(wallpaper.Height) * SCALE))

	return wallpaper
}

func (w *Wallpaper) Exec()  {
	err := w.getContent()
	if err != nil {
		fmt.Println("GET CONTENT ERROR", err)
		return
	}

	img, err := w.draw()
	if err != nil {
		fmt.Println("DRAW ERROR: ", err)
		return
	}

	SaveBaseOnOs(img)

	err = system.SetAsWallpaper(FilePath)
	if err != nil {
		fmt.Println("SET AS WALLPAPER ERROR: ", err)
		return
	}
	fmt.Println("Done")
}

//get earth content from gitee
func (w *Wallpaper) getContent() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return utils.BuildErrMsg("new request error", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return utils.BuildErrMsg("do request error", err)
	}

	////backup the http response for comparing with last wallpaper content
	w.Content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

//draw a wallpaper which size the save as screen resolution
func (w *Wallpaper) draw() (image.Image, error) {

	topLeft := image.Point{0, 0}
	BottomRight := image.Point{X: w.Width, Y: w.Height}

	//draw a new black background image which size the save as screen
	img := image.NewRGBA(image.Rectangle{topLeft, BottomRight})
	for x := 0; x < system.ScreenW; x++ {
		for y := 0; y < system.ScreenH; y++ {
			img.Set(x, y, color.Black)
		}
	}

	//calculate the side length base on screen resolution
	content, _ := png.Decode(bytes.NewReader(w.Content))

	//do resize
	content = resize.Resize(uint(w.ContentLength), uint(w.ContentLength), content, resize.Lanczos3)

	//calculate offset
	offset := image.Pt((w.Width-w.ContentLength)/2, (w.Height-w.ContentLength)/2)

	//do offset
	draw.Draw(img, content.Bounds().Add(offset), content, image.Point{}, draw.Over)

	return img, nil
}

//if os is windows,save the image as bmp file, otherwise png instead
func SaveBaseOnOs(img image.Image) error {
	if runtime.GOOS == "windows" {
		FilePath = os.TempDir() + string(os.PathSeparator) + "image.bmp"
		//create a new file for save wallpaper
		wpFile, _ := os.Create(FilePath)
		defer wpFile.Close()

		err := bmp.Encode(wpFile, img)
		if err != nil {
			return errors.New("windows decode image failed")
		}
		wpFile.Close()
	}else {
		FilePath = os.TempDir() + string(os.PathSeparator) + "image.png"
		wpFile, _ := os.Create(FilePath)
		defer wpFile.Close()

		err := png.Encode(wpFile, img)
		if err != nil {
			return errors.New("linux decode image failed")
		}
		wpFile.Close()
	}
	if CACHE {
		defer func() {
			os.MkdirAll(cacheDir, 0777)
			f,err := os.Create(cacheDir + string(os.PathSeparator) + time.Now().Format("20060102150405") + ".png")
			if err != nil {
				fmt.Println(err)
			}
			png.Encode(f, img)
			f.Close()
		}()
	}
	return nil
}


