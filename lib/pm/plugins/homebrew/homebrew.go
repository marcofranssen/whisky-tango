// +build darwin

package homebrew

// Homebrew installer to install apps via Homebrew
type Homebrew struct{}

// Install installs the given applications
func (i *Homebrew) Install(apps []string) error {
	err := Install(false, brews(apps))
	if err != nil {
		return err
	}
	err = Install(true, casks(apps))
	return err
}

// List lists applications installed with brew or brew cask
func (i *Homebrew) List() ([]string, error) {
	brews, err := List()
	if err != nil {
		return make([]string, 0), err
	}
	casks, err := ListCasks()
	if err != nil {
		return make([]string, 0), err
	}

	apps := make([]string, len(brews)+len(casks))
	for i, brew := range brews {
		apps[i] = string(brew)
	}
	fromBrews := len(brews) - 1
	for i, cask := range casks {
		apps[fromBrews+i] = string(cask)
	}

	return apps, nil
}
