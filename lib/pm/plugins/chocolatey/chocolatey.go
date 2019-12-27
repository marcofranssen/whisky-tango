// +build windows

package chocolatey

import "fmt"

// Chocolatey installer to install apps via Chocolatey
type Chocolatey struct{}

// Install installs the given applications
func (i *Chocolatey) Install(apps []string) error {
	return fmt.Errorf("chocolatey installer is not yet implemented")
}
