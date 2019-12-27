// +build darwin

package homebrew

import (
	"fmt"
	"strings"
)

// Homebrew installer to install apps via Homebrew
type Homebrew struct{}

// Install installs the given applications
func (i *Homebrew) Install(apps []string) error {
	fmt.Printf("  brew install %s\n", strings.Join(apps, " "))
	return nil
}

// List lists applications installed with brew or brew cask
func (i *Homebrew) List() ([]string, error) {
	fmt.Println("  brew list")
	fmt.Println("  brew cask list")

	return make([]string, 0), nil
}
