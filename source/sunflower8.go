package source

type Sunflower struct {
	Base
}

func (s *Sunflower) GetImageUrl() error {
	s.ImageUrl = "https://gitee.com/wencochen/No.8-Sunflower/raw/master/earth.png"
	return nil
}
