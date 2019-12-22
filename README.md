# 我的地球

让卫星给你拍壁纸

<img src="https://gitee.com/wencochen/No.8-Sunflower/raw/master/earth.png" width="100px" height="100px">

## 简介

此软件用于将向日葵8号气象卫星拍摄的地球影像的国内镜像设置为电脑桌面。

## 特点

1、 速度快

软件拉取的影像源放在国内，无不可抗力干扰，高效快速。

影像镜像仓库地址：[https://gitee.com/wencochen/No.8-Sunflower](https://gitee.com/wencochen/No.8-Sunflower)

2、 跨平台

本软件不区分32/64版本，可在Windows/Linux下使用。

目前已在win10、win7、ubuntu、deeping上进行过测试。

3、屏幕自适应，黄金分割

壁纸本身是通过代码画出来的，壁纸尺寸读取的是电脑分辨率，因此可完美适应任意电脑，壁纸内容高度为屏幕高度的0.618，舒适美观。

4、开机自启

软件可手动设置开机自启，将exe程序放置于`C:\ProgramData\Microsoft\Windows\Start Menu\Programs\StartUp`即可在系统启动时自动启动应用。

5、内存占用小

程序基于Golang编写，执行速度快，内存占用小，后台运行资源消耗可忽略不计。

## 编译配置

- SCALE 壁纸内容占比，默认 0.618

- CYCLE 更新周期，单位分钟，默认 10

- CACHE 是否缓存，默认 false，如开机则会在用户临时文件夹生成一个wallpaper_cache文件夹，用于保存历史壁纸。

## 编译

```base
$ go build
```

在windows下运行不出现DOS界面：

```base
$ go build -ldflags "-H windowsgui"
```

## 贡献

欢迎在各平台测试，提交pr。

