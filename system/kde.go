//+build linux

package system

import (
	"os/exec"
	"strconv"
)

func setKDEBackground(uri string) error {
	return exec.Command("qdbus", "org.kde.plasmashell", "/PlasmaShell", "org.kde.PlasmaShell.evaluateScript", `
		const monitors = desktops()
		for (var i = 0; i < monitors.length; i++) {
			monitors[i].currentConfigGroup = ["Wallpaper"]
			monitors[i].writeConfig("Image", `+strconv.Quote(uri)+`)
		}
	`).Run()
}
