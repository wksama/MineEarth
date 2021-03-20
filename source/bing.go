package source

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Bing struct {
	Base
}

func (b *Bing) GetImageUrl() error {
	resp, err := http.Get("https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-CN")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var bodyJson map[string]interface{}
	err = json.Unmarshal(bodyBytes, &bodyJson)
	if err != nil {
		return err
	}
	imageUrl := bodyJson["images"].([]interface{})[0].(map[string]interface{})["url"].(string)
	b.ImageUrl = "https://cn.bing.com/" + imageUrl
	return nil
}

func (b *Bing) ResizeImage() {

}

func (b *Bing) GetCanvas() {

}

func (b *Bing) Draw() {
	b.WallPaper = b.ImageContent
}

func (b *Bing) GetDuration() time.Duration {
	return 0
}
