package homebrew

import (
	"os/exec"
	"strings"
)

type App string
type AppList []App

func List() (AppList, error) {
	out, err := runBrewCmd("list", "-1")
	if err != nil {
		return make(AppList, 0), err
	}
	lines := strings.Split(out, "\n")

	apps := make(AppList, len(lines))
	for i, line := range lines {
		apps[i] = App(strings.Trim(line, " "))
	}
	return apps, nil
}

func ListCasks() (AppList, error) {
	out, err := runBrewCmd("cask", "list", "-1")
	if err != nil {
		return make(AppList, 0), err
	}
	lines := strings.Split(out, "\n")

	apps := make(AppList, len(lines))
	for i, line := range lines {
		apps[i] = App(strings.Trim(line, " "))
	}

	return apps, nil
}

func runBrewCmd(args ...string) (string, error) {
	cmd := exec.Command("brew", args...)
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}
