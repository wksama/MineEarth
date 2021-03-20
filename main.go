package main

import (
	"MineEarth/controller"
	"MineEarth/core"
	"MineEarth/source"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	core.AppInit()
	r := gin.Default()
	go source.RunWorkflow()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if viper.GetInt("source") != source.Scustomize {
			go source.RunWorkflow()
		}
	})

	r.LoadHTMLGlob("./template/*")
	r.StaticFS("/public", http.Dir("public"))

	r.GET("/", controller.Home)
	r.GET("/setting", controller.Setting)
	r.GET("/about", controller.About)

	r.POST("/saveSetting", controller.SaveSetting)

	_ = r.Run(":6280")
}
