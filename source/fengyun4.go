package source

type Fy struct {
	Base
}

func (f *Fy) GetImageUrl() error {
	f.ImageUrl = "http://img.nsmc.org.cn/CLOUDIMAGE/FY4A/MTCC/FY4A_DISK.JPG"
	return nil
}
