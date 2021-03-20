//+build windows

package system

import (
	"github.com/pkg/errors"
	"os"
	"path"
	"syscall"
	"unsafe"
)

var (
	w, _, _ = syscall.NewLazyDLL(`User32.dll`).NewProc(`GetSystemMetrics`).Call(uintptr(0))
	h, _, _ = syscall.NewLazyDLL(`User32.dll`).NewProc(`GetSystemMetrics`).Call(uintptr(1))
	ScreenW = int(w)
	ScreenH = int(h)
)

//将图片设为壁纸
func SetAsWallpaper(FilePath string) error {
	_, err := os.Stat(FilePath)
	if os.IsNotExist(err) {
		return errors.Wrap(err, "file not exit")
	}

	if path.Ext(FilePath) != ".bmp" {
		return errors.Wrap(nil, "file type must be bmp")
	}
	//set as wallpaper
	filenameUTF16, err := syscall.UTF16PtrFromString(FilePath)
	if err != nil {
		return errors.Wrap(err, "syscall.UTF16PtrFromString error")
	}

	user32 := syscall.NewLazyDLL("user32.dll")
	systemParametersInfo := user32.NewProc("SystemParametersInfoW")

	_, _, err = systemParametersInfo.Call(
		uintptr(0x0014), //20
		uintptr(0x0000), //0
		uintptr(unsafe.Pointer(filenameUTF16)),
		uintptr(0x01|0x02), //1|2
	)
	if err != nil && err.Error() != "The operation completed successfully." {
		return errors.Wrap(err, "systemParametersInfo.Call error")
	}
	return nil
}
