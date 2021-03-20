package core

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

var ConfigFilePath string
var CacheDir string

func AppInit() {
	log.SetFlags(log.Ldate | log.Ltime | log.LstdFlags | log.Llongfile)
	f, _ := os.OpenFile("./mineearth.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(f)

	configDir, _ := os.UserConfigDir()
	mineEarthConfigDir := configDir + string(os.PathSeparator) + "MineEarth"
	// config dir
	_ = os.MkdirAll(mineEarthConfigDir, 0666)
	ConfigFilePath = mineEarthConfigDir + string(os.PathSeparator) + "config.yaml"
	viper.SetConfigName("config")           // name of config file (without extension)
	viper.SetConfigType("yaml")             // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(mineEarthConfigDir) // path to look for the config file in
	_, err := os.Stat(ConfigFilePath)
	if os.IsNotExist(err) {
		var configExample = []byte(`
app:
  debug: false

wallpaper:
  source: 3
  saveName: "wallpaper"
  cache: false
  percentage: 0.618
`)
		err := viper.ReadConfig(bytes.NewBuffer(configExample))
		if err != nil {
			log.Fatal(err)
		}
		err = viper.WriteConfigAs(ConfigFilePath)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			log.Panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	userCacheDir, _ := os.UserCacheDir()
	CacheDir = userCacheDir + string(os.PathSeparator) + "MineEarth"
	// cache dir
	_ = os.MkdirAll(CacheDir, 0666)
}
