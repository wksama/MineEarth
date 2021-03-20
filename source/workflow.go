package source

import (
	"MineEarth/system"
	"bytes"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"image"
	"io/ioutil"
	"log"
	"mime/multipart"
	"reflect"
	"time"
)

var Ticker *time.Ticker
var itf WallpaperInterface

func RunWorkflow() {
	itf = GetStructByConfig()
	itf.Clean()
	err := Workflow(itf)
	if err != nil {
		log.Println("启动时设置壁纸错误：", err)
	}

	drt := itf.GetDuration()
	if drt == 0 {
		if Ticker != nil {
			Ticker.Stop()
		}
	} else {
		if Ticker == nil {
			Ticker = time.NewTicker(drt)
			for true {
				<-Ticker.C
				err := Workflow(itf)
				if err != nil {
					log.Println("定时设置壁纸错误：", err)
				}
			}
		} else {
			Ticker.Reset(drt)
		}
	}
}

func GetStructByConfig() WallpaperInterface {
	sourceType := viper.GetInt("wallpaper.source")
	sct := reflect.ValueOf(StructMap[sourceType]).Type()
	v := reflect.New(sct).Interface()
	var i WallpaperInterface
	i = v.(WallpaperInterface)

	return i
}

func Workflow(i WallpaperInterface) error {
	err := i.GetImageUrl()
	if err != nil {
	}
	err = i.GetImageContent()
	if err != nil {
		return errors.Wrap(err, "获取壁纸连接失败")
	}
	i.ResizeImage()
	i.GetCanvas()
	i.Draw()
	if i.GetWallpaper() != nil {
		path, err := i.SaveToFile()
		if err != nil {
			return errors.Wrap(err, "保存壁纸失败")
		}
		err = system.SetAsWallpaper(path)
		if err != nil {
			return errors.Wrap(err, "设置为壁纸失败")
		}
	}

	return nil
}

func ParseUploadedFile(fileHeader *multipart.FileHeader) (image.Image, error) {
	f, err := fileHeader.Open()
	if err != nil {
		return nil, errors.Wrap(err, "打开文件失败")
	}
	defer f.Close()

	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "读取文件失败")
	}
	img, _, err := image.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		return nil, errors.Wrap(err, "解析文件失败")
	}

	return img, nil
}
