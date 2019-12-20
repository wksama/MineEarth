// +build darwin

package system

import (
	"os/exec"
	"strconv"
)

var (
	ScreenW int
	ScreenH int
)

func init() {

}

// SetFromFile uses AppleScript to tell Finder to set the desktop wallpaper to specified file.
func SetAsWallpaper(file string) error {
	return exec.Command("osascript", "-e", `tell application "System Events" to tell every desktop to set picture to `+strconv.Quote(file)).Run()
}