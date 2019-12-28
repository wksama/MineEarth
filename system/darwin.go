// +build darwin

package system

import (
	// #cgo LDFLAGS: -framework CoreGraphics
	// #cgo LDFLAGS: -framework CoreFoundation
	// #include <CoreGraphics/CoreGraphics.h>
	// #include <CoreFoundation/CoreFoundation.h>
	"C"
	"os/exec"
	"strconv"
)

var (
	ScreenW int
	ScreenH int
)

func init() {
	displayID := C.CGMainDisplayID()
	ScreenW = int(C.CGDisplayPixelsWide(displayID))
	ScreenH = int(C.CGDisplayPixelsHigh(displayID))
}

// SetFromFile uses AppleScript to tell Finder to set the desktop wallpaper to specified file.
func SetAsWallpaper(file string) error {
	return exec.Command("osascript", "-e", `tell application "System Events" to tell every desktop to set picture to `+strconv.Quote(file)).Run()
}