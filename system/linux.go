//+build linux

package system

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)
var (
	Desktop = os.Getenv("XDG_CURRENT_DESKTOP")
	ErrUnsupportedDE = errors.New("your desktop environment is not supported")
	ScreenW int
	ScreenH int
)

// init guesses the current desktop by reading processes if $XDG_CURRENT_DESKTOP was not set.
func init() {
	command := "xdpyinfo | awk '/dimensions/{print $2}'"
	cmd := exec.Command("/bin/bash", "-c", command)
	opt,_ := cmd.Output()
	optArr := strings.Split(string(opt), "x")
	ScreenW,_ = strconv.Atoi(optArr[0])
	ScreenH,_ = strconv.Atoi(strings.TrimSpace(optArr[1]))

	if Desktop != "" {
		return
	}

	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		return
	}

	for _, file := range files {
		// continue if not pid
		_, err := strconv.ParseUint(file.Name(), 10, 64)
		if err != nil {
			continue
		}

		// checks to see if process's binary is a recognized window manager
		bin, err := os.Readlink("/proc/" + file.Name() + "/exe")
		if err != nil {
			continue
		}

		switch path.Base(bin) {
		case "i3":
			Desktop = "i3"
			return
		}
	}
}

// SetFromFile sets wallpaper from a file path.
func SetAsWallpaper(file string) error {
	if isGNOMECompliant() {
		return exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", strconv.Quote("file://"+file)).Run()
	}

	switch Desktop {
	case "KDE":
		return setKDEBackground("file://" + file)
	case "X-Cinnamon":
		return exec.Command("dconf", "write", "/org/cinnamon/desktop/background/picture-uri", strconv.Quote("file://"+file)).Run()
	case "MATE":
		return exec.Command("dconf", "write", "/org/mate/desktop/background/picture-filename", strconv.Quote(file)).Run()
	case "XFCE":
		desktops, err := getXFCEDesktops()
		if err != nil {
			return err
		}
		for _, desktop := range desktops {
			err := exec.Command("xfconf-query", "--channel", "xfce4-desktop", "--property", desktop, "--set", file).Run()
			if err != nil {
				return err
			}
		}
		return nil
	case "LXDE":
		return exec.Command("pcmanfm", "-w", file).Run()
	case "Deepin":
		command := "gsettings set com.deepin.wrap.gnome.desktop.background picture-uri " + file
		return exec.Command("/bin/bash", "-c", command).Run()
	case "i3":
		return exec.Command("feh", "--bg-fill", file).Run()
	default:
		return ErrUnsupportedDE
	}
}

func isGNOMECompliant() bool {
	return strings.Contains(Desktop, "GNOME") || Desktop == "Unity" || Desktop == "Pantheon"
}

func getXFCEDesktops() ([]string, error) {
	output, err := exec.Command("xfconf-query", "--channel", "xfce4-desktop", "--list").Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.Trim(string(output), "\n"), "\n")

	i := 0
	for _, line := range lines {
		if path.Base(line) == "last-image" {
			lines[i] = line
			i++
		}
	}
	lines = lines[:i]

	return lines, nil
}