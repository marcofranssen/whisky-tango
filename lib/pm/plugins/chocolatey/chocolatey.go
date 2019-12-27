// +build windows

package chocolatey

import "fmt"

// Chocolatey installer to install apps via Chocolatey
type Chocolatey struct{}

// Install installs the given applications
func (i *Chocolatey) Install(apps []string) error {
	return fmt.Errorf("chocolatey installer is not yet implemented")
}

// List lists applications installed with chocolatey
func (i *Homebrew) List() ([]string, error) {
	return make([]string, 0), fmt.Errorf("chocolatey app lister is not yet implemented")
}
