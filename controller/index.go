package controller

import (
	"MineEarth/source"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "golang.org/x/image/bmp"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
)

func Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.gohtml", gin.H{
		"path": "home",
	})
}
func Setting(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "setting.gohtml", gin.H{
		"settings": viper.AllSettings(),
		"choices":  source.NameMap,
		"path":     "setting",
	})
}

func About(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "about.gohtml", gin.H{
		"path": "about",
	})
}

type SettingForm struct {
	SaveName   string  `form:"save_name" binding:"required"`
	Source     int     `form:"source"`
	Cache      bool    `form:"cache"`
	Percentage float64 `form:"percentage"`
}

func SaveSetting(ctx *gin.Context) {
	var sf SettingForm
	err := ctx.ShouldBind(&sf)
	if err == nil {
		viper.Set("wallpaper.saveName", sf.SaveName)
		viper.Set("wallpaper.source", sf.Source)
		viper.Set("wallpaper.cache", sf.Cache)
		viper.Set("wallpaper.percentage", sf.Percentage)
		err = viper.WriteConfig()
		if err != nil {
			ctx.HTML(http.StatusBadRequest, "error.gohtml", "写如配置文件失败")
			return
		}
		if sf.Source == source.Scustomize {
			uploadedFile, err := ctx.FormFile("file")
			if err != nil {
				ctx.HTML(http.StatusBadRequest, "error.gohtml", err.Error())
				return
			}
			img, err := source.ParseUploadedFile(uploadedFile)
			if err != nil {
				ctx.HTML(http.StatusBadRequest, "error.gohtml", err.Error())
				return
			}
			sc := new(source.Customize)
			sc.SetWallpaper(img)
			err = source.Workflow(sc)
			if err != nil {
				ctx.HTML(http.StatusBadRequest, "error.gohtml", err.Error())
				return
			}
		}

		ctx.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	fmt.Println(err)
}
