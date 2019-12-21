package main

import (
	"os"
	"time"
)

const (
	URL   = "https://gitee.com/wencochen/No.8-Sunflower/raw/master/earth.png"
	SCALE = 0.618
	CYCLE = 10
	CACHE = false
)

var (
	FilePath  string
	cacheDir = os.TempDir() + string(os.PathSeparator) + "wallpaper_cache"
)

func main() {
	wallpaper := NewWallpaper()

	wallpaper.Exec()
	ticker := time.NewTicker(CYCLE * time.Minute)
	for  {
		<-ticker.C
		wallpaper.Exec()
	}
}