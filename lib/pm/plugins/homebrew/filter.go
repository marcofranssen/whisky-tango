package homebrew

import "strings"

func brews(apps []string) []string {
	brews := make([]string, 0)
	for _, app := range apps {
		if !strings.HasPrefix(app, "cask/") {
			brews = append(brews, app)
		}
	}
	return brews
}

func casks(apps []string) []string {
	casks := make([]string, 0)
	for _, app := range apps {
		if strings.HasPrefix(app, "cask/") {
			casks = append(casks, strings.Replace(app, "cask/", "", 1))
		}
	}
	return casks
}
